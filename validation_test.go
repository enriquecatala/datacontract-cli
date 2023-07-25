package main

import (
	"testing"
)

func TestValidateDataContractSpecification_HappyCase(t *testing.T) {
	e := ValidateDataContractSpecification("0.0.1")
	assertPassed(t, e)
}

func TestValidateDataContractSpecification_Illegal(t *testing.T) {
	e := ValidateDataContractSpecification("0.0.2")
	assertFailed(t, e, FieldNameDataContractSpecification, ValidationErrorReasonIllegalValue)
}

func TestValidateDataContractSpecification_Empty(t *testing.T) {
	e := ValidateDataContractSpecification("")
	assertFailed(t, e, FieldNameDataContractSpecification, ValidationErrorReasonEmptyString)
}

func TestValidateInfoId_HappyCase(t *testing.T) {
	e := ValidateInfoId("some-id")
	assertPassed(t, e)
}

func TestValidateInfoId_Empty(t *testing.T) {
	e := ValidateInfoId("")
	assertFailed(t, e, FieldNameInfoId, ValidationErrorReasonEmptyString)
}

func TestValidateInfoStartDate_HappyCase(t *testing.T) {
	date := "2021-12-08"
	e := ValidateInfoStartDate(&date)
	assertPassed(t, e)
}

func TestValidateInfoStartDate_Nil(t *testing.T) {
	e := ValidateInfoStartDate(nil)
	assertPassed(t, e)
}

func TestValidateInfoStartDate_IllegalFormat(t *testing.T) {
	date := "2021-December-08"
	e := ValidateInfoStartDate(&date)
	assertFailed(t, e, FieldNameInfoStartDate, ValidationErrorReasonIllegalFormat)
}

func TestValidateInfoEndDate_HappyCase(t *testing.T) {
	date := "2021-12-08"
	e := ValidateInfoEndDate(&date)
	assertPassed(t, e)
}

func TestValidateInfoEndDate_Nil(t *testing.T) {
	e := ValidateInfoEndDate(nil)
	assertPassed(t, e)
}

func TestValidateInfoEndDate_IllegalFormat(t *testing.T) {
	date := "2021-December-08"
	e := ValidateInfoEndDate(&date)
	assertFailed(t, e, FieldNameInfoEndDate, ValidationErrorReasonIllegalFormat)
}

func TestValidateProviderTeamId_HappyCase(t *testing.T) {
	e := ValidateProviderTeamId("some id")
	assertPassed(t, e)
}

func TestValidateProviderTeamId_Empty(t *testing.T) {
	e := ValidateProviderTeamId("")
	assertFailed(t, e, FieldNameProviderTeamId, ValidationErrorReasonEmptyString)
}

func TestValidateProviderDataProductId_HappyCase(t *testing.T) {
	e := ValidateProviderDataProductId("some id")
	assertPassed(t, e)
}

func TestValidateProviderDataProductId_Empty(t *testing.T) {
	e := ValidateProviderDataProductId("")
	assertFailed(t, e, FieldNameProviderDataProductId, ValidationErrorReasonEmptyString)
}

func TestValidateProviderOutputPortId_HappyCase(t *testing.T) {
	e := ValidateProviderOutputPortId("some id")
	assertPassed(t, e)
}

func TestValidateProviderOutputPortId_Empty(t *testing.T) {
	e := ValidateProviderOutputPortId("")
	assertFailed(t, e, FieldNameProviderOutputPortId, ValidationErrorReasonEmptyString)
}

func assertPassed(t *testing.T, e *ValidationError) {
	if e != nil {
		t.Error("Must pass.")
	}
}

func TestValidateConsumerTeamId_HappyCase(t *testing.T) {
	e := ValidateConsumerTeamId("some id")
	assertPassed(t, e)
}

func TestValidateConsumerTeamId_Empty(t *testing.T) {
	e := ValidateConsumerTeamId("")
	assertFailed(t, e, FieldNameConsumerTeamId, ValidationErrorReasonEmptyString)
}

func assertFailed(t *testing.T, e *ValidationError, field FieldName, reason ValidationErrorReason) {
	if e == nil {
		t.Error("Must fail.")
	}
	if e == nil || e.Field != field {
		t.Errorf("Field name must be '%v'", field)
	}
	if e == nil || e.Reason != reason {
		t.Errorf("Validation error reason must be '%v'", reason)
	}
}