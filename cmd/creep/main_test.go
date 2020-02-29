package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName = "creep"
)

func TestMain(m *testing.M) {
	fmt.Printf("Building %s...\n", binName)

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to build tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	if err := os.Remove(binName); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to remove tool %s: %s", binName, err)
		os.Exit(1)
	}

	images, err := filepath.Glob("*.jpg")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to find downloaded image files: %s", err)
	}
	if len(images) > 0 {
		for _, img := range images {
			if err := os.Remove(img); err != nil {
				fmt.Fprintf(os.Stderr, "Failed to remove image file %s: %s", img, err)
				os.Exit(1)
			}
		}
	}
	os.Exit(result)
}

func TestCreepCLI(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("ErrOnNoArgs", func(t *testing.T) {
		cmd := exec.Command(cmdPath)
		if err := cmd.Run(); err == nil {
			t.Fatal("Expected to receive error, none received")
		}
	})

	t.Run("PrintHelp", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-h")
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("PrintVersion", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-v")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		exp := "Master" + "\n"
		if exp != string(out) {
			t.Errorf("Expected %q, got %q instead\n", exp, string(out))
		}
	})

	t.Run("DownloadSingleImage", func(t *testing.T) {
		args := "-u https://source.unsplash.com/random"
		cmd := exec.Command(cmdPath, strings.Split(args, " ")...)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DownloadMultipleImages", func(t *testing.T) {
		args := "-u https://source.unsplash.com/random -c 2"
		cmd := exec.Command(cmdPath, strings.Split(args, " ")...)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DownloadSingleImageWithName", func(t *testing.T) {
		args := "-u https://source.unsplash.com/random -n test"
		cmd := exec.Command(cmdPath, strings.Split(args, " ")...)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DownloadSingleImageWithOutpath", func(t *testing.T) {
		args := "-u https://source.unsplash.com/random -o test"
		cmd := exec.Command(cmdPath, strings.Split(args, " ")...)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
		if err := os.RemoveAll("test"); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to remove test directory: %q", err)
			os.Exit(1)
		}
	})

	t.Run("DownloadMultipleImagesWithThrottle", func(t *testing.T) {
		args := "-u https://source.unsplash.com/random -c 2 -t 4"
		cmd := exec.Command(cmdPath, strings.Split(args, " ")...)
		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})
}
