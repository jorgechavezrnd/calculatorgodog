package calc_test

import (
	calc "calculatorgodog"
	"fmt"

	"github.com/cucumber/godog"
)

type CalcScenarioContext struct {
	*godog.ScenarioContext
	calc *calc.Calculator
}

func (ctx *CalcScenarioContext) calculatorIsCleared() error {
	ctx.calc.Clear()
	return nil
}

func (ctx *CalcScenarioContext) iPress(num int) error {
	ctx.calc.Press(num)
	return nil
}

func (ctx *CalcScenarioContext) iAdd(num int) error {
	ctx.calc.Add(num)
	return nil
}

func (ctx *CalcScenarioContext) iSubstract(num int) error {
	ctx.calc.Sub(num)
	return nil
}

func (ctx *CalcScenarioContext) iMultiplyBy(factor int) error {
	return ctx.calc.MultiplyBy(factor)
}

func (ctx *CalcScenarioContext) theResultShouldBe(expectedResult int) error {
	result := ctx.calc.Result()
	if result == expectedResult {
		return nil
	}
	return fmt.Errorf("%d doesn't match expectation: %d", result, expectedResult)
}

func FeatureContext(ctx *godog.ScenarioContext) {
	calcctx := CalcScenarioContext{
		calc:            new(calc.Calculator),
		ScenarioContext: ctx,
	}

	ctx.Step(`^calculator is cleared$`, calcctx.calculatorIsCleared)
	ctx.Step(`^i press (\d+)$`, calcctx.iPress)
	ctx.Step(`^i add (\d+)$`, calcctx.iAdd)
	ctx.Step(`^i substract (\d+)$`, calcctx.iSubstract)
	ctx.Step(`^i multiply by (\d+)$`, calcctx.iMultiplyBy)
	ctx.Step(`^the result should be (\d+)$`, calcctx.theResultShouldBe)
}
