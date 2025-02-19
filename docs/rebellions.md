# Rebellions
이 문서는 Accelerator Operator가  리벨리온 NPU를 어떻게 쿠버네티스 상에서 지원하는지 설명합니다.

## 워크로드 지원 방식
Accelerator Operator는 리벨리온 NPU를 쿠버네티스 상에서 두 가지 방식으로 지원합니다.

1. `VM Passthrough` - VM에 NPU를 직접 할당하여 사용합니다.
2. `Container` - 쿠버네티스 파드에 할당하여 사용합니다.

## 소프트웨어 컴포넌트 구성
각각의 방식에 필요한 소프트웨어 컴포넌트는 다음과 같습니다.

1. VM에 NPU Passthrough하여 할당:
   - `VFIO Manager` - 노드에 존재하는 npu에 vfio-pci 드라이버를 바인드하기 위함
   
2. 컨테이너로 할당
   - `Rebellions Driver Manager` - 리벨리온 드라이버 설치를 위함
   - `Rebellions K8S Device Plugin` - 쿠버네티스 상에 NPU 리소스를 인식시키고 사용하기 위함
   - `Metric Exporter` - NPU 모니터링을 위함

## 컴포넌트 관리 방식
Accelerator Operator는 각 노드의 `rebellions.ai/npu.workload.config` 레이블을 참조하여 데몬셋 기반의 컴포넌트들을 관리합니다.  
레이블에 정의된 워크로드 유형에 따라 Accelerater Operator는 아래와 같은 노드 레이블을 추가합니다.
- `rebellions.ai/npu.workload.config: vm-passthrough`일 시
    ```yaml
    "rebellions.ai/npu.deploy.vfio-manager":          "true",
    "rebellions.ai/npu.deploy.npu-feature-discovery": "true",
    ```
- `rebellions.ai/npu.workload.config: container`일 시
    ```yaml
    "rebellions.ai/npu.deploy.driver":                "true",
    "rebellions.ai/npu.deploy.device-plugin":         "true",
    "rebellions.ai/npu.deploy.metric-exporter":       "true",
    "rebellions.ai/npu.deploy.npu-feature-discovery": "true",
    ```
각 컴포넌트들의 디폴트 매니페스트는 `assets/rebellions/` 디렉토리에 정의되어있으며, `AcceleratorConfig` CR의 `.spec.providers.rebellions`를 통해 세부 스펙을 설정할 수 있습니다.  
설정 가능한 스펙은 다음과 같습니다.
- `RebellionsConfigSpec` - 리벨리온 컴포넌트 스펙들의 집합입니다.
    ```go
    type RebellionsConfigSpec struct {
        // Daemonsets defines common configurations for rebellions daemonsets
        Daemonsets *DaemonsetsSpec `json:"daemonsets,omitempty"`
        // Driver spec
        Driver *RebellionsDriverSpec `json:"driver,omitempty"`
        // DevicePlugin spec
        DevicePlugin *RebellionsDevicePluginSpec `json:"devicePlugin,omitempty"`
        // MetricExporter spec
        MetricExporter *RebellionsMetricExporterSpec `json:"metricExporter,omitempty"`
        // NPUFeatureDiscovery spec
        NPUFeatureDiscovery *RebellionsNPUFeatureDiscoverySpec `json:"nfd,omitempty"`
        // VFIOManager spec
        VFIOManager *VFIOManagerSpec `json:"vfioManager,omitempty"`
    }
    ```
- `DaemonsetSpec` - 리벨리온 컴포넌트들의 공통된 데몬셋 설정을 지원합니다.
    ```go
    type DaemonsetsSpec struct {
        // Set labels
        Labels map[string]string `json:"labels,omitempty"`
        // Set annotations
        Annotations map[string]string `json:"annotations,omitempty"`
        // Set tolerations
        Tolerations []corev1.Toleration `json:"tolerations,omitempty"`
        // +kubebuilder:validation:Optional
        PriorityClassName string `json:"priorityClassName,omitempty"`
    }
    ```
- `VFIOManagerSpec` - vfio-manager 컴포넌트 세부 설정에 관한 스펙입니다.
    ```go
    type VFIOManagerSpec struct {
        // Flag which indicates deploying VFIO Manager with operator is enabled
        Enabled bool `json:"enabled"`
        // VFIO manager container image
        Image *ImageSpec `json:"image,omitempty"`
        // container resource spec
        Resources *ResourceRequirements `json:"resources,omitempty"`
        // additional container environments
        Env []EnvVar `json:"env,omitempty"`
    }
    ```