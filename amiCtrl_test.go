package main

import (
    "bytes"
    "os/exec"
    _ "fmt"
    "strings"
    _ "time"
    "testing"
)

func TestVersionFlag(t *testing.T) {
    cmd := exec.Command("gom", "run", "amiCtrl.go", "-version")
    stdout := new(bytes.Buffer)
    cmd.Stdout = stdout

    _ = cmd.Run()

    if ! strings.Contains(stdout.String(), AppVersion) {
        t.Fatal("Failed Test")
    }
}

func TestStdoutList(t *testing.T) {
    cmd := exec.Command("sh", "tests/test_stdout_list.sh")
    stdout := new(bytes.Buffer)
    cmd.Stdout = stdout
    output := "test-image"

    _ = cmd.Run()

    if ! strings.Contains(stdout.String(), output) {
        t.Fatal("Failed Test")
    }
}

func TestStdoutCreate(t *testing.T) {
    cmd := exec.Command("sh", "tests/test_stdout_create.sh")
    stdout := new(bytes.Buffer)
    cmd.Stdout = stdout
    output := "test-image99999"

    _ = cmd.Run()

    if ! strings.Contains(stdout.String(), output) {
        t.Fatal("Failed Test")
    }
}

func TestStdoutCreateError(t *testing.T) {
    cmd := exec.Command("sh", "tests/test_stdout_create_error.sh")
    stdout := new(bytes.Buffer)
    cmd.Stdout = stdout
    output := "`-instance` オプションを指定して下さい."

    _ = cmd.Run()

    if ! strings.Contains(stdout.String(), output) {
        t.Fatal("Failed Test")
    }
}

func TestStdoutDelete(t *testing.T) {
    cmd := exec.Command("sh", "tests/test_stdout_delete.sh")
    stdout := new(bytes.Buffer)
    cmd.Stdout = stdout
    output := "AMI を削除しました."

    _ = cmd.Run()

    if ! strings.Contains(stdout.String(), output) {
        t.Fatal("Failed Test")
    }
}

func TestStdoutDeleteError(t *testing.T) {
    cmd := exec.Command("sh", "tests/test_stdout_delete_error.sh")
    stdout := new(bytes.Buffer)
    cmd.Stdout = stdout
    output := "`-ami` オプションを指定して下さい."

    _ = cmd.Run()

    if ! strings.Contains(stdout.String(), output) {
        t.Fatal("Failed Test")
    }
}

func TestStdoutDeleteNo(t *testing.T) {
    cmd := exec.Command("sh", "tests/test_stdout_delete_no.sh")
    stdout := new(bytes.Buffer)
    cmd.Stdout = stdout
    output := "処理を停止します."

    _ = cmd.Run()

    if ! strings.Contains(stdout.String(), output) {
        t.Fatal("Failed Test")
    }
}

func TestStdoutState(t *testing.T) {
    cmd := exec.Command("sh", "tests/test_stdout_state.sh")
    stdout := new(bytes.Buffer)
    cmd.Stdout = stdout
    output := "test-image88888"

    _ = cmd.Run()

    if ! strings.Contains(stdout.String(), output) {
        t.Fatal("Failed Test")
    }
}

/*
func TestConvertDate(t *testing.T) {
    str := "2018-09-28 22:52:24 +0000 UTC"
    layout := "2006-01-02 15:04:05 +0000 UTC"
    tm, _ := time.Parse(layout, str)

    actual := convertDate(tm)
    expected := "2018-09-29 07:52:24"

    if actual != expected {
        t.Fatal("Failed Test")
    }
}
*/