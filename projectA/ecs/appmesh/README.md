
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


## Virtual Gateway Path 주의!!

Virtual Gateway의 Route 설정에서 prefix match를 설정하면 virtual gateway를 거쳐가면서 prefix 경로가 벗겨집니다.

예를 들어, `/chapter` prefix match를 사용하고 /chapter/1 요청이 들어올 경우 virtual gateway 뒷단의 어플리케이션에서는 /1 로 요청을 받습니다.

- REF : https://docs.aws.amazon.com/ko_kr/app-mesh/latest/userguide/gateway-routes.html