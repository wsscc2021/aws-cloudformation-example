
## 작업 순서 주의!!

1. CloudMap

   - Private Namespace 생성

   - Service 생성

2. appmesh 생성

   - Virtual Node 생성 (우선 Backend 없이 생성합니다.)

   - Virtual Router 생성

   - Virtual Service 생성

   - Virtual Node의 Backend 설정

3. ecs service 생성

## 권한 주의!!

- arn:aws:iam::aws:policy/AWSAppMeshEnvoyAccess

- arn:aws:iam::aws:policy/CloudWatchFullAccess

- arn:aws:iam::aws:policy/AWSXRayDaemonWriteAccess

## Security Group

AppMesh를 사용하면 Envoy Proxy 를 거쳐 통신하지만, Envoy Proxy 포트를 별도로 열어줄 필요가 없습니다.

단지 **어플리케이션의 포트만 열어주면 됩니다.**

