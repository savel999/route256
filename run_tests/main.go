package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const outBinaryName = "__app_test"
const goPath = "/usr/lib/go-1.22/bin/go"

type TaskTest struct {
	TestStdinPath  string
	TestResultPath string
}

func main() {
	ctx := context.Background()
	taskDir := os.Args[1]

	//скомпилить бинарник
	_ = os.Remove(taskDir + "/" + outBinaryName)

	buildCmd := exec.CommandContext(
		ctx,
		goPath,
		"build",
		"-o",
		taskDir+"/"+outBinaryName,
		taskDir,
	)

	if err := buildCmd.Run(); err != nil {
		log.Fatal(fmt.Errorf("build error: %w", err))
	}

	tests, err := parseTests(taskDir + "/tests")
	if err != nil {
		log.Fatal(fmt.Errorf("не удалось распарсить тесты задачи: %w", err))
	}

	log.Printf("Необходимо выполнить тестов: %d шт.\n", len(tests))

	for _, test := range tests {
		log.Printf("Запуск теста: %s.", test.TestStdinPath)

		commandText := fmt.Sprintf(
			"%s/%s < %s > %s.res",
			taskDir, outBinaryName, test.TestStdinPath, test.TestStdinPath,
		)

		execTestCmd := exec.CommandContext(ctx, "bash", "-c", commandText)
		execTestCmd.Stdout = os.Stdout
		execTestCmd.Stderr = os.Stderr

		log.Printf("Текст команды: %s", execTestCmd.String())

		start := time.Now()
		err = execTestCmd.Run()
		log.Printf("Время выполнения: %d мс", time.Since(start).Milliseconds())

		if err != nil {
			log.Fatal(fmt.Errorf("execute error: %w", err))
		}

		diffTestCmd := exec.CommandContext(
			ctx,
			"diff",
			test.TestResultPath,
			test.TestStdinPath+".res",
		)

		diffOutput, _ := diffTestCmd.CombinedOutput()

		log.Printf("Дифы тестов: %s\n", string(diffOutput))

		break
	}

	fmt.Println(tests)
	fmt.Println(err)
}

func parseTests(dir string) ([]TaskTest, error) {
	dirFiles, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	resultTemplatePath, result := regexp.MustCompile(`\d+\.a`), []TaskTest{}

	for _, file := range dirFiles {
		if file.IsDir() {
			continue
		}

		if resultTemplatePath.MatchString(file.Name()) {
			result = append(result, TaskTest{
				TestStdinPath:  filepath.Join(dir, strings.ReplaceAll(file.Name(), ".a", "")),
				TestResultPath: filepath.Join(dir, file.Name()),
			})
		}
	}

	return result, nil
}
