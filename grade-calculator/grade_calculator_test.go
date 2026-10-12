package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment",70, Assignment)
	gradeCalculator.AddGrade("exam 1", 71, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 72, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 61, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 62, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 51, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 4, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
    if Assignment.String() != "assignment" {
        t.Error("Expected assignment")
    }

    if Exam.String() != "exam" {
        t.Error("Expected exam")
    }

    if Essay.String() != "essay" {
        t.Error("Expected essay")
    }
}

func TestEmptyGrades(t *testing.T) {
    gradeCalculator := NewGradeCalculator()

    actual_value := gradeCalculator.GetFinalGrade()
    if actual_value != "F"{
        t.Errorf("Expected F for empty grades; got %s", actual_value)
    }
}

func TestPass(t *testing.T) {
	gradeCalculator := NewGradeCalculator(true)

	gradeCalculator.AddGrade("assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam", 100, Exam)
	gradeCalculator.AddGrade("essay", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if actual_value != "Pass" {
		t.Errorf("Expected Pass; got %s", actual_value)
	}
}

func testFail(t *testing.T) {
	gradeCalculator := NewGradeCalculator(true)

	gradeCalculator.AddGrade("assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam", 50, Exam)
	gradeCalculator.AddGrade("essay", 50, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if actual_value != "Fail" {
		t.Errorf("Expected Fail; got %s", actual_value)
	}
}