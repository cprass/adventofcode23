package utils

type Runner interface {
	Run(input []string, isPartOne bool) (string, error)
}
