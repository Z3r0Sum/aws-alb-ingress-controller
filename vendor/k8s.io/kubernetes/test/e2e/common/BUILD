package(default_visibility = ["//visibility:public"])

licenses(["notice"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = [
        "apparmor.go",
        "autoscaling_utils.go",
        "configmap.go",
        "container_probe.go",
        "docker_containers.go",
        "downward_api.go",
        "downwardapi_volume.go",
        "empty_dir.go",
        "events.go",
        "expansion.go",
        "host_path.go",
        "init_container.go",
        "kubelet_etc_hosts.go",
        "networking.go",
        "pods.go",
        "privileged.go",
        "projected.go",
        "secrets.go",
        "sysctl.go",
        "util.go",
        "volumes.go",
    ],
    tags = ["automanaged"],
    deps = [
        "//pkg/api:go_default_library",
        "//pkg/api/v1:go_default_library",
        "//pkg/api/v1/pod:go_default_library",
        "//pkg/apis/autoscaling/v1:go_default_library",
        "//pkg/apis/extensions/v1beta1:go_default_library",
        "//pkg/client/clientset_generated/clientset:go_default_library",
        "//pkg/client/clientset_generated/internalclientset:go_default_library",
        "//pkg/client/conditions:go_default_library",
        "//pkg/kubelet:go_default_library",
        "//pkg/kubelet/sysctl:go_default_library",
        "//pkg/security/apparmor:go_default_library",
        "//test/e2e/framework:go_default_library",
        "//test/utils:go_default_library",
        "//vendor:github.com/golang/glog",
        "//vendor:github.com/onsi/ginkgo",
        "//vendor:github.com/onsi/gomega",
        "//vendor:golang.org/x/net/websocket",
        "//vendor:k8s.io/apimachinery/pkg/api/errors",
        "//vendor:k8s.io/apimachinery/pkg/api/resource",
        "//vendor:k8s.io/apimachinery/pkg/apis/meta/v1",
        "//vendor:k8s.io/apimachinery/pkg/fields",
        "//vendor:k8s.io/apimachinery/pkg/labels",
        "//vendor:k8s.io/apimachinery/pkg/runtime",
        "//vendor:k8s.io/apimachinery/pkg/util/intstr",
        "//vendor:k8s.io/apimachinery/pkg/util/sets",
        "//vendor:k8s.io/apimachinery/pkg/util/uuid",
        "//vendor:k8s.io/apimachinery/pkg/util/wait",
        "//vendor:k8s.io/apimachinery/pkg/watch",
        "//vendor:k8s.io/client-go/tools/cache",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)
