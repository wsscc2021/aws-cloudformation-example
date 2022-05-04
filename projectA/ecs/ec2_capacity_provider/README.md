
## Healthcheck endpoint 주의!!

ECS 앞단에 ELB를 붙일 경우 Healthcheck 포인트가 2개이므로 이를 조정하는 작업에 주의합니다.

일반적으로 Healthcheck 엔드포인트를 같게 설정하는 것이 좋습니다.

- **ELB TargetGroup**의 HealthCheck Path

- **ECS Task Definition**의 HealthCheck Command



## 작업 순서 주의!!

리소스 **생성** 시에는 아래와 같은 순서로 작업을 진행합니다.

1. EC2 Auto Scaling Group 생성

2. ECS Cluster 생성

3. ELB & Target Group 생성

4. (콘솔 작업) ELB TargetGroup을 Listener Rule로 추가

5. ECR Repository 생성 및 이미지 푸쉬

6. ECS Service 생성


리소스 **삭제** 시에는 아래와 같은 순서로 작업을 진행합니다.

1. ECS Service 삭제

2. ECR Repository 삭제

3. (콘솔 작업) ELB TargetGroup을 Listener Rule에서 제거

4. ELB & Target Group 삭제

5. **EC2 Auto Scaling Group 삭제**

6. ECS Cluster 삭제



## Cloudwatch Log Group 주의!!

Cloudwatch Log Group이 이미 생성되어 있는 경우 ecs service 생성에 오류가 발생합니다.

만약 생성되어 있다면 삭제한 뒤 ecs service를 생성합니다.

기존 Log Group을 삭제하지 못하는 경우, ecs service template을 적절히 수정합니다.


## 인스턴스 수준 지표 모니터링

ECS Cluster에 Container Insight를 활성화하면 클러스터 및 서비스 수준 지표만 수집됩니다.

즉, 인스턴스 수준 지표는 수집되지 않습니다. 따라서, CloudWatch Agent를 Daemon Service로 

배포하는 작업을 별도로 수행해야합니다.

관련 링크 : 

- https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/deploy-container-insights-ECS-instancelevel.html

- https://ecsworkshop.com/monitoring/container_insights/setup/#enable-instance-level-insights