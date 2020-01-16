// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "mdtogo"; DO NOT EDIT.
package fn

var READMEShort = `Run local containers against Resource configuration`
var READMELong = `
Functions are executables packaged in container images which accept a collection of
Resource configuration as input, and emit a collection of Resource configuration as output.

They may be used to:

- Generate configuration from Templates, DSLs, CRD-style abstractions, key-value pairs, etc. -- e.g.
  expand Helm charts, JSonnet, etc.
- Inject fields or otherwise modifying configuration -- e.g. add init-containers, side-cars, etc
- Rollout configuration changes across an organization -- e.g. similar to
  https://github.com/reactjs/react-codemod
- Validate configuration -- e.g. ensure Organizational policies are enforced

Functions may be run either imperatively with ` + "`" + `kpt run DIR/ --image` + "`" + ` or declaratively with
` + "`" + `kpt run DIR/` + "`" + ` and specifying them in files.

Functions specified in files must contain an annotation to mark them as function declarations:

      annotations:
        config.kubernetes.io/function: |
          container:
            image: gcr.io/example.com/image:version
        config.kubernetes.io/local-config: "true"

Functions may be run at different times depending on the function and the organizational needs:

- as part of the build and development process
- as pre-commit checks
- as PR checks
- as pre-release checks
- as pre-rollout checks
`
var READMEExamples = `
    kpt fn run DIR/ --image gcr.io/example.com/my-fn

    kpt fn run DIR/ --fn-path FUNCTIONS_DIR/`

var RunShort = `Execute functional programs from container images to generate, modify or validate configuration.`
var RunLong = `
` + "`" + `run` + "`" + ` sequentially locally executes one or more programs which may modify Resource configuration.

**Architecture:**

- Programs are packaged as container images which are pulled and run locally
- Input Resource configuration is read from some source and written to container STDIN
- Output Resource configuration is read from container STDOUT and written to some sink
- If the container exits non-0, fail with an error message

**Caveats:**

- If ` + "`" + `DIR` + "`" + ` is provided, it is used as both the *source* and *sink*.
- A function may be explicitly specified with ` + "`" + `--image` + "`" + `

  Example: Locally run the container image gcr.io/example.com/my-fn against the Resources
  in DIR/.


        kpt fn run DIR/ --image gcr.io/example.com/my-fn

- If not ` + "`" + `DIR` + "`" + ` is specified, the source will default to STDIN and sink will default to STDOUT

  Example: This is equivalent to the preceding example.


        kpt source DIR/ | kpt fn run --image gcr.io/example.com/my-fn | kpt sink DIR/

- Arguments specified after ` + "`" + `--` + "`" + ` will be provided to the function as a ConfigMap input 

  Example: In addition to the input Resources, provide to the container image a ConfigMap
  containing ` + "`" + `data: {foo: bar}` + "`" + `.  This is used to set the behavior of the function.


        # run the my-fn image, configured with foo=bar
        kpt fn run DIR/ --image gcr.io/example.com/my-fn:v1.0.0 -- foo=bar

- Alternatively functions and their input configuration may be declared in
  files rather than directly on the command line
- ` + "`" + `FUNCTIONS_DIR` + "`" + ` may optionally be under the Resource ` + "`" + `DIR` + "`" + `

  Example: This is equivalent to the preceding example.
  Rather than specifying ` + "`" + `gcr.io/example.com/my-fn` + "`" + ` as a flag, specify it in a file using the
  ` + "`" + `config.kubernetes.io/function` + "`" + ` annotation, and discover it with ` + "`" + `--fn-path` + "`" + `.


        # run the my-fn, configured with foo=bar -- fn is declared in a file
        kpt fn run DIR/ --fn-path FUNCTIONS_DIR/

        # FUNCTIONS_DIRS/some.yaml
        apiVersion: v1
        kind: ConfigMap
        metdata:
          annotations:
            config.kubernetes.io/function: |
              container:
                image: gcr.io/example.com/my-fn
        data:
          foo: bar

- Additionally, functions may be discovered implicitly by putting them in ` + "`" + `run` + "`" + ` *source*.


        # run the my-fn, configured with foo=bar -- fn is declared in the input
        kpt fn run DIR/

        # DIR/functions/some.yaml
        apiVersion: v1
        kind: ConfigMap
        metdata:
          annotations:
            config.kubernetes.io/function: |
              container:
                image: gcr.io/example.com/my-fn
        data:
          foo: bar

- Functions which are nested under some sub directory are scoped only to Resources under that
  same sub directory.  This allows fine grain control over how functions are executed.


         Example: gcr.io/example.com/my-fn is scoped to Resources in stuff/ and
         is NOT scoped to Resources in apps/
             .
             ├── stuff
             │   ├── inscope-deployment.yaml
             │   ├── stuff2
             │   │     └── inscope-deployment.yaml
             │   └── functions # functions in this dir are scoped to stuff/...
             │       └── some.yaml
             └── apps
                 ├── not-inscope-deployment.yaml
                 └── not-inscope-service.yaml

- Multiple functions may be specified.  If they are specified in the same file they will
  be run in the same order that they are specified.
- Functions may define their own API input types - these may be client-side equivalents of CRDs.


        kpt fn run DIR/

        # DIR/functions/some.yaml
        apiVersion: v1
        kind: ConfigMap
        metdata:
          annotations:
            config.kubernetes.io/function: |
              container:
                image: gcr.io/example.com/my-fn
        data:
          foo: bar
        ---
        apiVersion: v1
        kind: MyType
        metdata:
          annotations:
            config.kubernetes.io/function: |
              container:
                image: gcr.io/example.com/other-fn
        spec:
          field:
            nestedField: value

#### Arguments:

  DIR:
    Path to local directory.

#### Config Functions:

  Config functions are specified as Kubernetes types containing a metadata.annotations.[config.kubernetes.io/function]
  field specifying an image for the container to run.  This image tells run how to invoke the container.

  Example config function:

	# in file example/fn.yaml
	apiVersion: fn.example.com/v1beta1
	kind: ExampleFunctionKind
	metadata:
	  annotations:
	    config.kubernetes.io/function: |
	      container:
	        # function is invoked as a container running this image
	        image: gcr.io/example/examplefunction:v1.0.1
	    config.kubernetes.io/local-config: "true" # tools should ignore this
	spec:
	  configField: configValue

  In the preceding example, 'kpt cfg run example/' would identify the function by
  the metadata.annotations.[config.kubernetes.io/function] field.  It would then write all Resources in the directory to
  a container stdin (running the gcr.io/example/examplefunction:v1.0.1 image).  It
  would then write the container stdout back to example/, replacing the directory
  file contents.
`
var RunExamples = `
    # read the Resources from DIR, provide them to a container my-fun as input,
    # write my-fn output back to DIR
    kpt fn run DIR/ --image gcr.io/example.com/my-fn

    # provide the my-fn with an input ConfigMap containing ` + "`" + `data: {foo: bar}` + "`" + `
    kpt fn run DIR/ --image gcr.io/example.com/my-fn:v1.0.0 -- foo=bar

    # run the functions in FUNCTIONS_DIR against the Resources in DIR
    kpt fn run DIR/ --fn-path FUNCTIONS_DIR/

    # discover functions in DIR and run them against Resource in DIR.
    # functions may be scoped to a subset of Resources -- see ` + "`" + `kpt help fn run` + "`" + `
    kpt fn run DIR/`

var SinkShort = `Implement a Sink by writing input to a local directory.`
var SinkLong = `
Implement a Sink by writing input to a local directory.

    kpt fn sink DIR

  DIR:
    Path to local directory.

` + "`" + `sink` + "`" + ` writes its input to a directory
`
var SinkExamples = `
    kpt fn source DIR/ | your-function | kpt fn sink DIR/`

var SourceShort = `Implement a Source by reading a local directory.`
var SourceLong = `
Implement a Source by reading a local directory.

    kpt fn source DIR

  DIR:
    Path to local directory.

` + "`" + `source` + "`" + ` emits configuration to act as input to a function
`
var SourceExamples = `
    # emity configuration directory as input source to a function
    kpt fn source DIR/

    kpt fn source DIR/ | your-function | kpt fn sink DIR/`