build:
	cd account_service && docker build -t account_service:latest .;
	cd insurance_service && docker build -t insurance_service:latest .;
	cd loan_service && docker build -t loan_service:latest .;

apply:
	kubectl apply -f account_service/deployment.yaml
	kubectl apply -f insurance_service/deployment.yaml
	kubectl apply -f loan_service/deployment.yaml

apply-jaeger:
	helm repo add jaegertracing https://jaegertracing.github.io/helm-charts;
	helm upgrade -i jaeger jaegertracing/jaeger --values jaeger.yaml;

apply-otel-collector:
	helm repo add open-telemetry https://open-telemetry.github.io/opentelemetry-helm-charts;
	helm upgrade -i opentelemetry-collector open-telemetry/opentelemetry-collector --values otel.yaml;