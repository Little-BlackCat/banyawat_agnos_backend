package main

import "testing"

func TestCheckStrongPassword(t *testing.T) {
	// Test a strong password
	t.Run("StrongPassword", func(t *testing.T) {
		password := "1445D1cd"
		actions, isStrong := CheckStrongPassword(password)

		if actions != 0 || !isStrong {
			t.Errorf("Expected a strong password, but got actions=%d and isStrong=%v", actions, isStrong)
		}
	})

	// Test a weak password
	t.Run("WeakPassword", func(t *testing.T) {
		password := "Aa1"
		actions, isStrong := CheckStrongPassword(password)

		if actions == 0 || isStrong {
			t.Errorf("Expected a weak password, but got actions=%d and isStrong=%v", actions, isStrong)
		}
	})
}