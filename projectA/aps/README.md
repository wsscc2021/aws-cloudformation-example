
refer

- https://github.com/prometheus-community/helm-charts


## aps

- create aps on aws console

## prometheus

- create iam policy

  `ingest-proxy-policy.json`
  ```json
  {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Action": [
          "aps:RemoteWrite", 
          "aps:GetSeries", 
          "aps:GetLabels",
          "aps:GetMetricMetadata"
        ], 
        "Resource": "*"
      }
    ]
  }
  ```
  ```powershell
  aws iam create-policy \
    --policy-name ingest-proxy-policy
    --policy-document file://ingest-proxy-policy.json
  ```

- create service account for iam role
  ```
  eksctl eksctl create iamserviceaccount \
    --cluster=${CLUSTER_NAME} \
    --namespace=prometheus-system \
    --name=prometheus \
    --attach-policy-arn=arn:aws:iam::${ACCOUNT_ID}:policy/ingest-proxy-policy \
    --override-existing-serviceaccounts \
    --region ${REGION} \
    --approve
  ```

- add helm repo
  ```
  helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
  helm repo update
  ```

- install kube-state-metrics helm chart
  ```
  helm install kube-state-metrics prometheus-community/kube-state-metrics -n prometheus-system -f kube-state-metrics_override.yaml
  ```

- install prometheus helm chart
  ```
  helm install prometheus prometheus-community/prometheus -n prometheus-system -f prometheus_override.yaml
  ```