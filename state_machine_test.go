package nitpicking_test

import (
	"testing"

	"github.com/MarioCarrion/nitpicking"
)

func TestImportsSectionMachine_Transition(t *testing.T) {
	tests := [...]struct {
		name          string
		startState    nitpicking.ImportsSection
		validStates   []nitpicking.ImportsSection
		invalidStates []nitpicking.ImportsSection
	}{
		{
			"Standard",
			nitpicking.ImportsSectionStd,
			[]nitpicking.ImportsSection{
				nitpicking.ImportsSectionStd,
				nitpicking.ImportsSectionExternal,
				nitpicking.ImportsSectionLocal,
			},
			[]nitpicking.ImportsSection{},
		},
		{
			"External",
			nitpicking.ImportsSectionExternal,
			[]nitpicking.ImportsSection{
				nitpicking.ImportsSectionExternal,
				nitpicking.ImportsSectionLocal,
			},
			[]nitpicking.ImportsSection{nitpicking.ImportsSectionStd},
		},
		{
			"Local",
			nitpicking.ImportsSectionLocal,
			[]nitpicking.ImportsSection{nitpicking.ImportsSectionLocal},
			[]nitpicking.ImportsSection{
				nitpicking.ImportsSectionStd,
				nitpicking.ImportsSectionExternal,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(ts *testing.T) {
			for _, s := range tt.validStates {
				machine := nitpicking.NewImportsSectionMachine(tt.startState)
				if err := machine.Transition(s); err != nil {
					t.Fatalf("expected no error, to %s", err)
				}
			}

			for _, s := range tt.invalidStates {
				machine := nitpicking.NewImportsSectionMachine(tt.startState)
				if err := machine.Transition(s); err == nil {
					t.Fatalf("expected error, got nil")
				}
			}
		})
	}
}

func TestSectionMachine_Transition(t *testing.T) {
	tests := [...]struct {
		name          string
		factory       func() nitpicking.SectionMachine
		validStates   []nitpicking.Section
		invalidStates []nitpicking.Section
	}{
		{
			"Start",
			func() nitpicking.SectionMachine {
				return nitpicking.SectionMachine{}
			},
			[]nitpicking.Section{nitpicking.SectionImports},
			[]nitpicking.Section{
				nitpicking.SectionStart,
				nitpicking.SectionConsts,
				nitpicking.SectionTypes,
				nitpicking.SectionVars,
				nitpicking.SectionFuncs,
				nitpicking.SectionMethods,
			},
		},
		{
			"Imports",
			func() nitpicking.SectionMachine {
				v := nitpicking.SectionMachine{}
				v.Transition(nitpicking.SectionImports)
				return v
			},
			[]nitpicking.Section{nitpicking.SectionConsts},
			[]nitpicking.Section{
				nitpicking.SectionStart,
				nitpicking.SectionImports,
				nitpicking.SectionTypes,
				nitpicking.SectionVars,
				nitpicking.SectionFuncs,
				nitpicking.SectionMethods,
			},
		},
		{
			"Consts",
			func() nitpicking.SectionMachine {
				v := nitpicking.SectionMachine{}
				v.Transition(nitpicking.SectionImports)
				v.Transition(nitpicking.SectionConsts)
				return v
			},
			[]nitpicking.Section{
				nitpicking.SectionTypes,
				nitpicking.SectionConsts,
			},
			[]nitpicking.Section{
				nitpicking.SectionStart,
				nitpicking.SectionImports,
				nitpicking.SectionVars,
				nitpicking.SectionFuncs,
				nitpicking.SectionMethods,
			},
		},
		{
			"Types",
			func() nitpicking.SectionMachine {
				v := nitpicking.SectionMachine{}
				v.Transition(nitpicking.SectionImports)
				v.Transition(nitpicking.SectionConsts)
				v.Transition(nitpicking.SectionTypes)
				return v
			},
			[]nitpicking.Section{
				nitpicking.SectionVars,
			},
			[]nitpicking.Section{
				nitpicking.SectionStart,
				nitpicking.SectionImports,
				nitpicking.SectionTypes,
				nitpicking.SectionConsts,
				nitpicking.SectionFuncs,
				nitpicking.SectionMethods,
			},
		},
		{
			"Vars",
			func() nitpicking.SectionMachine {
				v := nitpicking.SectionMachine{}
				v.Transition(nitpicking.SectionImports)
				v.Transition(nitpicking.SectionConsts)
				v.Transition(nitpicking.SectionTypes)
				v.Transition(nitpicking.SectionVars)
				return v
			},
			[]nitpicking.Section{
				nitpicking.SectionFuncs,
			},
			[]nitpicking.Section{
				nitpicking.SectionStart,
				nitpicking.SectionImports,
				nitpicking.SectionConsts,
				nitpicking.SectionTypes,
				nitpicking.SectionVars,
				nitpicking.SectionMethods,
			},
		},
		{
			"Funcs",
			func() nitpicking.SectionMachine {
				v := nitpicking.SectionMachine{}
				v.Transition(nitpicking.SectionImports)
				v.Transition(nitpicking.SectionConsts)
				v.Transition(nitpicking.SectionTypes)
				v.Transition(nitpicking.SectionVars)
				v.Transition(nitpicking.SectionFuncs)
				return v
			},
			[]nitpicking.Section{
				nitpicking.SectionFuncs,
				nitpicking.SectionMethods,
			},
			[]nitpicking.Section{
				nitpicking.SectionStart,
				nitpicking.SectionImports,
				nitpicking.SectionConsts,
				nitpicking.SectionTypes,
				nitpicking.SectionVars,
			},
		},
		{
			"Methods",
			func() nitpicking.SectionMachine {
				v := nitpicking.SectionMachine{}
				v.Transition(nitpicking.SectionImports)
				v.Transition(nitpicking.SectionConsts)
				v.Transition(nitpicking.SectionTypes)
				v.Transition(nitpicking.SectionVars)
				v.Transition(nitpicking.SectionFuncs)
				v.Transition(nitpicking.SectionMethods)
				return v
			},
			[]nitpicking.Section{
				nitpicking.SectionMethods,
			},
			[]nitpicking.Section{
				nitpicking.SectionStart,
				nitpicking.SectionImports,
				nitpicking.SectionConsts,
				nitpicking.SectionTypes,
				nitpicking.SectionVars,
				nitpicking.SectionFuncs,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(ts *testing.T) {
			for _, s := range tt.validStates {
				validator := tt.factory()
				if err := validator.Transition(s); err != nil {
					t.Fatalf("expected no error, got %s", err)
				}
			}

			for _, s := range tt.invalidStates {
				validator := tt.factory()
				if err := validator.Transition(s); err == nil {
					t.Fatalf("expected error, got nil")
				}
			}
		})
	}
}