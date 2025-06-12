# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=dsa

# Main package path
MAIN_PACKAGE=./cmd

# Determine the operating system
ifeq ($(OS),Windows_NT)
    BINARY_NAME := $(BINARY_NAME).exe
    RM := del /Q
else
    RM := rm -f
endif

all: test build

build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PACKAGE)

test:
	$(GOTEST) -v ./...

test-stack:
	$(GOTEST) -v ./internal/stack

test-queue:
	$(GOTEST) -v ./internal/queue

test-linkedlist:
	$(GOTEST) -v ./internal/linkedlist

test-binarytree:
	$(GOTEST) -v ./internal/binarytree

test-avltree:
	$(GOTEST) -v ./internal/avltree

test-bubble:
	$(GOTEST) -v ./internal/bubblesort

test-selection:
	$(GOTEST) -v ./internal/selectionsort

test-neetcode150:
	$(GOTEST) -v ./internal/neetcode150

test-two-sum:
	$(GOTEST) -v ./internal/neetcode150 -run TestTwoSum

test-valid-anagram:
	$(GOTEST) -v ./internal/neetcode150 -run TestValidAnagram

test-two-sum-ii:
	$(GOTEST) -v ./internal/neetcode150 -run TestTwoSumII

test-majority-element:
	$(GOTEST) -v ./internal/neetcode150 -run TestMajorityElement

test-contains-duplicate-ii:
	$(GOTEST) -v ./internal/neetcode150 -run TestContainsDuplicateII

clean:
	$(GOCLEAN)
	$(RM) $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

deps:
	$(GOGET) ./...

tidy:
	$(GOMOD) tidy

.PHONY: all build test clean run deps tidy test-stack test-queue test-linkedlist test-binarytree test-avltree test-bubble test-selection test-neetcode150 test-two-sum test-valid-anagram test-two-sum-ii test-majority-element test-contains-duplicate-ii