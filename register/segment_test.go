package register

import (
	"github.com/aakash-rajur/http/params"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cleanPath(t *testing.T) {
	t.Parallel()

	type args struct {
		path string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test_cleanPath_1",
			args{
				path: "/",
			},
			"/",
		},
		{
			"test_cleanPath_2",
			args{
				path: "/test",
			},
			"/test",
		},
		{
			"test_cleanPath_3",
			args{
				path: "/test/",
			},
			"/test",
		},
		{
			"test_cleanPath_4",
			args{
				path: "/test//",
			},
			"/test",
		},
		{
			"test_cleanPath_5",
			args{
				path: "/test//test",
			},
			"/test/test",
		},
		{
			"test_cleanPath_6",
			args{
				path: "/test//test/",
			},
			"/test/test",
		},
		{
			"test_cleanPath_7",
			args{
				path: "/test//test//",
			},
			"/test/test",
		},
		{
			"test_cleanPath_8",
			args{
				path: "/test//test//test",
			},
			"/test/test/test",
		},
		{
			"test_cleanPath_9",
			args{
				path: "/test//test//test/",
			},
			"/test/test/test",
		},
		{
			"test_cleanPath_10",
			args{
				path: "/test//test//test//",
			},
			"/test/test/test",
		},
		{
			"test_cleanPath_11",
			args{
				path: "/test//test//test//test",
			},
			"/test/test/test/test",
		},
		{
			"test_cleanPath_12",
			args{
				path: "/test//test//test//test/",
			},
			"/test/test/test/test",
		},
		{
			"test_cleanPath_13",
			args{
				path: "/test//test//test//test//",
			},
			"/test/test/test/test",
		},
		{
			"test_cleanPath_14",
			args{
				path: "/test//test//test//test//test",
			},
			"/test/test/test/test/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cleanPath(tt.args.path)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_segment_isParam(t *testing.T) {
	t.Parallel()

	type fields struct {
		s string
	}

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"test_segment_isParam_1",
			fields{
				s: "{test}",
			},
			true,
		},
		{
			"test_segment_isParam_2",
			fields{
				s: "test",
			},
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := segment(tt.fields.s)
			got := s.isParam()

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_segment_name(t *testing.T) {
	t.Parallel()

	type fields struct {
		s string
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			"test_segment_name_1",
			fields{
				s: "{test}",
			},
			"test",
		},
		{
			"test_segment_name_2",
			fields{
				s: "test",
			},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := segment(tt.fields.s)
			got := s.name()

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_segment_cmp(t *testing.T) {
	t.Parallel()

	type args struct {
		other segment
	}

	tests := []struct {
		name string
		s    segment
		args args
		want int
	}{
		{
			"test_segment_cmp_1",
			segment("test"),
			args{
				other: segment("test"),
			},
			0,
		},
		{
			"test_segment_cmp_2",
			segment("test"),
			args{
				other: segment("{test}"),
			},
			1,
		},
		{
			"test_segment_cmp_3",
			segment("{test}"),
			args{
				other: segment("test"),
			},
			0,
		},
		{
			"test_segment_cmp_4",
			segment("{test}"),
			args{
				other: segment("test"),
			},
			0,
		},
		{
			"test_segment_cmp_5",
			segment("{test}"),
			args{
				other: segment("{test}"),
			},
			0,
		},
		{
			"test_segment_cmp_6",
			segment("{test}"),
			args{
				other: segment("{test2}"),
			},
			1,
		},
		{
			"test_segment_cmp_7",
			segment("{test2}"),
			args{
				other: segment("{test}"),
			},
			-1,
		},
		{
			"test_segment_cmp_8",
			segment("{test2}"),
			args{
				other: segment("{test2}"),
			},
			0,
		},
		{
			"test_segment_cmp_9",
			segment("{test2}"),
			args{
				other: segment("{test3}"),
			},
			-1,
		},
		{
			"test_segment_cmp_10",
			segment("{arg1}"),
			args{
				other: segment("test"),
			},
			0,
		},
		{
			"test_segment_cmp_11",
			segment("{arg1}"),
			args{
				other: segment("test"),
			},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.cmp(tt.args.other)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_segments_cmp(t *testing.T) {
	t.Parallel()

	type args struct {
		other  segments
		strict bool
	}

	tests := []struct {
		name string
		s    segments
		args args
		want int
	}{
		{
			name: "test_segments_cmp_1",
			s:    segments{},
			args: args{
				other:  segments{},
				strict: false,
			},
		},
		{
			name: "test_segments_cmp_2",
			s:    segments{},
			args: args{
				other:  segments{"test"},
				strict: false,
			},
			want: -1,
		},
		{
			name: "test_segments_cmp_3",
			s:    segments{"test"},
			args: args{
				other:  segments{},
				strict: false,
			},
			want: 1,
		},
		{
			name: "test_segments_cmp_4",
			s:    segments{"test"},
			args: args{
				other:  segments{"test"},
				strict: false,
			},
		},
		{
			name: "test_segments_cmp_5",
			s:    segments{"test"},
			args: args{
				other:  segments{"test", "test"},
				strict: false,
			},
			want: -1,
		},
		{
			name: "test_segments_cmp_6",
			s:    segments{"test", "test"},
			args: args{
				other:  segments{"test"},
				strict: false,
			},
			want: 1,
		},
		{
			name: "test_segments_cmp_7",
			s:    segments{"test", "test"},
			args: args{
				other:  segments{"test", "test"},
				strict: false,
			},
		},
		{
			name: "test_segments_cmp_8",
			s:    segments{"test", "test"},
			args: args{
				other:  segments{"test", "test", "test"},
				strict: false,
			},
			want: -1,
		},
		{
			name: "test_segments_cmp_9",
			s:    segments{"test", "test", "test"},
			args: args{
				other:  segments{"test", "test"},
				strict: false,
			},
			want: 1,
		},
		{
			name: "test_segments_cmp_10",
			s:    segments{"test", "test", "test"},
			args: args{
				other:  segments{"test", "test", "test"},
				strict: false,
			},
		},
		{
			name: "test_segments_cmp_11",
			s:    segments{"test", "{id}"},
			args: args{
				other:  segments{"test", "test", "test"},
				strict: false,
			},
			want: -1,
		},
		{
			name: "test_segments_cmp_12",
			s:    segments{"{arg1}", "test", "{arg2}"},
			args: args{
				other:  segments{"test", "test"},
				strict: false,
			},
			want: 1,
		},
		{
			name: "test_segments_cmp_13",
			s:    segments{"{arg1}", "{arg2}", "{arg3}"},
			args: args{
				other:  segments{"test", "test", "test"},
				strict: false,
			},
		},
		{
			name: "test_segments_cmp_14",
			s:    segments{"{arg1}", "{arg2}", "{arg3}"},
			args: args{
				other:  segments{"test", "test", "test", "test"},
				strict: false,
			},
			want: -1,
		},
		{
			name: "test_segments_cmp_15",
			s: segments{
				"PATCH",
				"repos",
				"{owner}",
				"{repo}",
				"pulls",
				"comments",
				"{number}",
			},
			args: args{
				other: segments{
					"PATCH",
					"repos",
					"{owner}",
					"{repo}",
					"pulls",
					"{number}",
				},
				strict: true,
			},
			want: 1,
		},
		{
			name: "test_segments_cmp_16",
			s: segments{
				"PATCH",
				"repos",
				"{owner}",
				"{repo}",
				"pulls",
				"{number}",
			},
			args: args{
				other: segments{
					"PATCH",
					"repos",
					"{owner}",
					"{repo}",
					"pulls",
					"comments",
					"{number}",
				},
				strict: true,
			},
			want: -1,
		},
		{
			name: "test_segments_cmp_16",
			s: segments{
				"GET",
				"repos",
				"{owner}",
				"{repo}",
				"git",
				"blobs",
				"{sha}",
			},
			args: args{
				other: segments{
					"GET",
					"repos",
					"{owner}",
					"{repo}",
					"downloads",
					"{id}",
				},
				strict: true,
			},
			want: 1,
		},
		{
			name: "test_segments_cmp_18",
			s: segments{
				"GET",
				"repos",
				"{owner}",
				"{repo}",
				"downloads",
				"{id}",
			},
			args: args{
				other: segments{
					"GET",
					"repos",
					"{owner}",
					"{repo}",
					"git",
					"blobs",
					"{sha}",
				},
				strict: true,
			},
			want: -1,
		},
		{
			name: "test_segments_cmp_19",
			s: segments{
				"GET",
				"repos",
				"{owner}",
				"{repo}",
				"forks",
			},
			args: args{
				other: segments{
					"GET",
					"repos",
					"{owner}",
					"{repo}",
					"downloads",
					"{id}",
				},
				strict: true,
			},
			want: -1,
		},
		{
			name: "test_segments_cmp_20",
			s: segments{
				"GET",
				"repos",
				"{owner}",
				"{repo}",
				"downloads",
				"{id}",
			},
			args: args{
				other: segments{
					"GET",
					"repos",
					"{owner}",
					"{repo}",
					"forks",
				},
				strict: true,
			},
			want: 1,
		},
		{
			name: "test_segments_cmp_21",
			s: segments{
				"PATCH",
				"repos",
				"{owner}",
				"{repo}",
				"pulls",
				"comments",
				"{number}",
			},
			args: args{
				other: segments{
					"PATCH",
					"repos",
					"{owner}",
					"{repo}",
					"pulls",
					"{number}",
				},
				strict: true,
			},
			want: 1,
		},
		{
			name: "test_segments_cmp_22",
			s: segments{
				"PATCH",
				"repos",
				"{owner}",
				"{repo}",
				"pulls",
				"{number}",
			},
			args: args{
				other: segments{
					"PATCH",
					"repos",
					"{owner}",
					"{repo}",
					"pulls",
					"comments",
					"{number}",
				},
				strict: true,
			},
			want: -1,
		},
		{
			name: "test_segments_cmp_23",
			s: segments{
				"PATCH",
				"repos",
				"{owner}",
				"{repo}",
				"pulls",
				"comments",
				"{number}",
			},
			args: args{
				other: segments{
					"PATCH",
					"repos",
					"{owner}",
					"{repo}",
					"pulls",
					"{number}",
				},
				strict: true,
			},
			want: 1,
		},
		{
			name: "test_segments_cmp_24",
			s: segments{
				"PATCH",
				"repos",
				"{owner}",
				"{repo}",
				"pulls",
				"{number}",
			},
			args: args{
				other: segments{
					"PATCH",
					"repos",
					"{owner}",
					"{repo}",
					"pulls",
					"comments",
					"{number}",
				},
				strict: true,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.cmp(tt.args.other)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_segments_params(t *testing.T) {
	t.Parallel()

	type args struct {
		other segments
	}

	tests := []struct {
		name string
		s    segments
		args args
		want params.Params
	}{
		{
			name: "test_segments_params_1",
			s:    segments{},
			args: args{
				other: segments{},
			},
			want: params.Params{},
		},
		{
			name: "test_segments_params_2",
			s:    segments{},
			args: args{
				other: segments{"test"},
			},
			want: params.Params{},
		},
		{
			name: "test_segments_params_3",
			s:    segments{"test"},
			args: args{
				other: segments{"test"},
			},
			want: params.Params{},
		},
		{
			name: "test_segments_params_4",
			s:    segments{"{arg1}", "{arg2}", "{arg3}"},
			args: args{
				other: segments{"test", "test", "test"},
			},
			want: params.Params{
				"arg1": "test",
				"arg2": "test",
				"arg3": "test",
			},
		},
		{
			name: "test_segments_params_5",
			s:    segments{"{arg1}", "{arg2}", "{arg3}"},
			args: args{
				other: segments{"test", "test", "test", "test"},
			},
			want: params.Params{
				"arg1": "test",
				"arg2": "test",
				"arg3": "test",
			},
		},
		{
			name: "test_segments_params_6",
			s:    segments{"{arg1}", "{arg2}", "{arg3}"},
			args: args{
				other: segments{"test", "test"},
			},
			want: params.Params{
				"arg1": "test",
				"arg2": "test",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.params(tt.args.other)

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_segmentsFromPath(t *testing.T) {
	t.Parallel()

	type args struct {
		pattern string
	}

	tests := []struct {
		name string
		args args
		want segments
	}{
		{
			name: "test_segmentsFromPath_1",
			args: args{
				pattern: "/",
			},
			want: segments{""},
		},
		{
			name: "test_segmentsFromPath_2",
			args: args{
				pattern: "/test",
			},
			want: segments{"test"},
		},
		{
			name: "test_segmentsFromPath_3",
			args: args{
				pattern: "/test/",
			},
			want: segments{"test"},
		},
		{
			name: "test_segmentsFromPath_4",
			args: args{
				pattern: "/test//",
			},
			want: segments{"test"},
		},
		{
			name: "test_segmentsFromPath_5",
			args: args{
				pattern: "/test//test",
			},
			want: segments{"test", "test"},
		},
		{
			name: "test_segmentsFromPath_6",
			args: args{
				pattern: "/test//test/",
			},
			want: segments{"test", "test"},
		},
		{
			name: "test_segmentsFromPath_7",
			args: args{
				pattern: "/test//test//",
			},
			want: segments{"test", "test"},
		},
		{
			name: "test_segmentsFromPath_8",
			args: args{
				pattern: "/test//test//test",
			},
			want: segments{"test", "test", "test"},
		},
		{
			name: "test_segmentsFromPath_9",
			args: args{
				pattern: "/test//test//test/",
			},
			want: segments{"test", "test", "test"},
		},
		{
			name: "test_segmentsFromPath_10",
			args: args{
				pattern: "/test//test//test//",
			},
			want: segments{"test", "test", "test"},
		},
		{
			name: "test_segmentsFromPath_11",
			args: args{
				pattern: "/test//test//test//test",
			},
			want: segments{"test", "test", "test", "test"},
		},
		{
			name: "test_segmentsFromPath_12",
			args: args{
				pattern: "/test//test//test//test/",
			},
			want: segments{"test", "test", "test", "test"},
		},
		{
			name: "test_segmentsFromPath_13",
			args: args{
				pattern: "/:arg1/:arg2/:arg3/:arg4",
			},
			want: segments{":arg1", ":arg2", ":arg3", ":arg4"},
		},
		{
			name: "test_segmentsFromPath_14",
			args: args{
				pattern: "/:arg1/:arg2/:arg3/:arg4/",
			},
			want: segments{":arg1", ":arg2", ":arg3", ":arg4"},
		},
		{
			name: "test_segmentsFromPath_15",
			args: args{
				pattern: "/:arg1/:arg2/:arg3/:arg4//",
			},
			want: segments{":arg1", ":arg2", ":arg3", ":arg4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := segmentsFromPath(tt.args.pattern)

			assert.Equal(t, tt.want, got)
		})
	}
}
