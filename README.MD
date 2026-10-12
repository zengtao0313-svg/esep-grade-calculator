
# ESEP Grade Calculator

This is a simple grade calculator written in Go. It calculates final letter grades using assignments, exams, and essays.

## Requirements

- Golang 1.24+

## Project Files

- `go.mod` - Go module information
- `grade_calculator.go` - Grade calculator code
- `grade_calculator_test.go` - Unit tests

## Grade Weights

- Assignments: 50%
- Exams: 35%
- Essays: 15%

The calculator uses the weighted average to determine the final grade.

## Grade Calculation

The final grade is calculated using a weighted average:

Final Grade = (Assignment Average×0.50) + (Exam Average×0.35) + (Essay Average×0.15)

Each category average is calculated by adding all grades in that category and dividing by the number of grades.

## Letter Grades

- A: 90-100
- B: 80-89
- C: 70-79
- D: 60-69
- F: Below 60

A grade of C or higher is passing.

## Running Tests

From the repository root, run:

```bash
cd grade-calculator
go test -v ./...
```

## Checking Code Coverage

From the `grade-calculator` directory, run:

```bash
go test -coverprofile=coverage.txt ./...
go tool cover -func=coverage.txt
```

## Example

```go
calculator := NewGradeCalculator()

calculator.AddGrade("Assignment 1", 90, Assignment)
calculator.AddGrade("Exam 1", 85, Exam)
calculator.AddGrade("Essay 1", 95, Essay)

finalGrade := calculator.GetFinalGrade()
```

The original project contains some code and test issues that will be fixed later.
