# Accelerator Config

## Intro
`AcceleratorConfig`는 GPU, NPU를 비롯한 AI 가속기를 쿠버네티스 상에서 지원하기 위한 컴포넌트 구성을 정의하는 CRD(Custom Resource Definition)입니다.

## Reconcile 로직
`AcceleratorConfig`의 reconcile 로직은 가속기의 자체 오퍼레이터 제공 범위에 따라 크게 두가지 방식으로 나뉩니다.
1. 가속기 자체 오퍼레이터를 통한 컴포넌트 관리
   - `AcceleratorConfig` 스펙에 따라 가속기 오퍼레이터를 배포 및 관리합니다.
   - `AcceleratorConfig` 스펙에 따라 해당 오퍼레이터의 커스텀 리소스를 배포합니다.
2. Accelerator Operator를 통한 컴포넌트 관리
   - `AcceleratorConfig` 스펙에 따라 assets에 정의된 리소스를 업데이트합니다.
   - 업데이트 된 리소스를 배포합니다.

**대략적인 리컨사일 로직은 아래 그림과 같습니다.**  
<img width="600" alt="Image" src="https://github.com/user-attachments/assets/4b3853c5-fb27-47f0-9e46-8a724e958fdf" />