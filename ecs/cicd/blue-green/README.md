
## PreRequisite

- Application Load Balancer

- ECS Cluster & Service (CODE_DEPLOY 형태로 생성)

- Artifact Bucket

## Point of modification

`taskdef.json`
- 미리 생성한 ECS Service의 TaskDefinition을 복사해온 뒤 수정합니다.
- `TaskDefinitionArn` , `requiresAttributes` , `revision` , `status` 와 같은 필요없는 설정은 삭제합니다.
- `image` 를 `<IMAGE1_NAME>` 으로 변경합니다.
  ```
  "image": "<IMAGE1_NAME>",
  ```

`appspec.yaml`
- 컨테이너 이름과 포트번호를 수정합니다.

`buildspec.yml`
- 이미지 태그의 규칙이 지정되어 있을 경우 수정합니다.

`DeploymentConfig`
- `DeploymentConfig` 리소스를 수정하면 Blue/Green배포 뿐만 아니라 Canaray 배포, Linear 배포도 가능합니다.
- canary 예시
  ```
  #
  # CodeDeploy - Deployment Config
  #
  DeploymentConfig:
    Type: AWS::CodeDeploy::DeploymentConfig
    Properties: 
      ComputePlatform: ECS # ECS | Lambda | Server
      DeploymentConfigName: 'ecs-canary10'
      TrafficRoutingConfig: 
        TimeBasedCanary: 
          CanaryInterval: 5 # minute 0 , 5 , 10 ...
          CanaryPercentage: 10
        # TimeBasedLinear: 
        #   LinearInterval: 5 # minute 0 , 5 , 10 ...
        #   LinearPercentage: 25
        Type: TimeBasedCanary # AllAtOnce | TimeBasedCanary | TimeBasedLinear
  ```
- linear 예시
  ```
  #
  # CodeDeploy - Deployment Config
  #
  DeploymentConfig:
    Type: AWS::CodeDeploy::DeploymentConfig
    Properties: 
      ComputePlatform: ECS # ECS | Lambda | Server
      DeploymentConfigName: 'ecs-linear25'
      TrafficRoutingConfig: 
        # TimeBasedCanary: 
        #   CanaryInterval: 5 # minute 0 , 5 , 10 ...
        #   CanaryPercentage: 10
        TimeBasedLinear: 
          LinearInterval: 5 # minute 0 , 5 , 10 ...
          LinearPercentage: 25
        Type: TimeBasedLinear # AllAtOnce | TimeBasedCanary | TimeBasedLinear
  ```