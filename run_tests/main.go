package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const outBinaryName = "__app_test"
const goPath = "/snap/go/10679/bin/go"

type TaskTest struct {
	TestStdinPath  string
	TestResultPath string
	TestNumber     int
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

	var totalExecutionTime int64 = 0

	for _, test := range tests {
		commandText := fmt.Sprintf(
			"%s/%s < %s > %s.res",
			taskDir, outBinaryName, test.TestStdinPath, test.TestStdinPath,
		)

		execTestCmd := exec.CommandContext(ctx, "bash", "-c", commandText)
		execTestCmd.Stdout = os.Stdout
		execTestCmd.Stderr = os.Stderr

		start := time.Now()
		err = execTestCmd.Run()
		executionDurationInMs := time.Since(start).Milliseconds()

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

		testStatus := "❌"

		if len(diffOutput) == 0 {
			testStatus = "✅"
		}

		log.Printf("%s  Тест %d: %d мс", testStatus, test.TestNumber, executionDurationInMs)

		totalExecutionTime += executionDurationInMs
	}

	log.Printf("Суммарное время выполнения: %d мс", totalExecutionTime)
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
			taskTest := TaskTest{
				TestStdinPath:  filepath.Join(dir, strings.ReplaceAll(file.Name(), ".a", "")),
				TestResultPath: filepath.Join(dir, file.Name()),
			}

			pathParts := strings.Split(taskTest.TestStdinPath, "/")
			testNumber, _ := strconv.Atoi(pathParts[len(pathParts)-1])

			taskTest.TestNumber = testNumber

			result = append(result, taskTest)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].TestNumber < result[j].TestNumber
	})

	return result, nil
}
