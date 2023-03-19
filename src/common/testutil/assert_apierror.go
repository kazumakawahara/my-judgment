package testutil

import (
	"reflect"
	"testing"

	"my-judgment/common/apperr"
	"my-judgment/common/mjerr"
)

func AssertAPIError(t *testing.T, want, got error) {
	t.Helper()

	if want == nil && got == nil {
		return
	}

	wantOriginErr := asOriginError(t, want)
	gotOriginErr := asOriginError(t, got)

	if !reflect.DeepEqual(wantOriginErr, gotOriginErr) {
		wantAPIErr := apperr.AsAPIError(wantOriginErr)
		if wantAPIErr == nil {
			t.Fatalf("failed in type assertion for want apiError: %+v", want)
		}

		gotAPIErr := apperr.AsAPIError(gotOriginErr)
		if gotAPIErr == nil {
			t.Fatalf("failed in type assertion for got apiError: %+v", got)
		}

		t.Errorf(
			"differs: (-wantErr +gotErr)\n- %s\n  %+v\n+ %s\n  %+v",
			wantAPIErr.Error(),
			wantAPIErr.Detail(),
			gotAPIErr.Error(),
			gotAPIErr.Detail(),
		)
	}
}

func asOriginError(t *testing.T, err error) mjerr.OriginError {
	t.Helper()

	mjErr := mjerr.AsApoError(err)
	if mjErr == nil {
		t.Fatalf("failed in type assertion for mjError: %+v", err)
	}

	appErr := mjErr.OriginError()
	if appErr == nil {
		t.Fatalf("failed in type assertion for OriginError: %+v", err)
	}

	return appErr
}
