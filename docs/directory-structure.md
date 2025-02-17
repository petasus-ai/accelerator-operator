# 디렉토리 구성
```
accelerator-operator/
├── api/                                            # API 관련 코드            
│   └── petasus/                                    #
│       └── v1beta1/                                #
│           └── acceleratornodeconfig_types.go      #
├── assets/                                         # 오퍼레이터가 관리할 에셋 매니페스트
│   ├── vfio-manager/                               #
│   └── rebellions/                                 # 리벨리온 관련 에셋
│       ├── node-feature-discovery/                 #
│       ├── device-plugin/                          #
│       ├── driver-manager/                         #
│       └── metric-exporter/                        #
├── build/                                          # 빌드 관련 파일
│   └── Dockerfile                                  #
├── cmd/                                            # 애플리케이션 진입점
│   └── main.go                                     #
├── config/                                         # 오퍼레이터 Kustomize 구성 파일 및 설정
├── controllers/                                    # 컨트롤러 로직
│   └── accleratornodeconfig_controller.go          #
├── deployments/                                    # Helm 배포 관련 파일
│   └── accelerator-operator/                       #
├── docs/                                           # 문서화 관련 파일
├── hack/                                           # 개발 및 테스트 스크립트
├── internal/                                       # 내부 패키지
├── scripts/                                        # 유틸리티 스크립트
├── go.mod                                          # Go 모듈 파일
├── go.sum                                          # Go 모듈 의존성 파일
└── Makefile                                        # 빌드 및 배포 스크립트
```
