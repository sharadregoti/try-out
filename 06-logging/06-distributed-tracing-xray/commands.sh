public.ecr.aws/aws-observability/aws-otel-collector:latest

docker run -p 4317:4317  -p 4318:4318  -p 55679:55679  -v ${pwd}/config/otel.yaml:/otel.yaml  --env AWS_REGION="ap-south-1" --env AWS_ACCESS_KEY_ID="" --env AWS_SECRET_ACCESS_KEY=""  public.ecr.aws/aws-observability/aws-otel-collector:latest --config /otel.yaml

docker run --rm --name fluent-bit -p 3000:3000 -v $(pwd)/config/fluent-bit.conf:/fluent-bit/etc/fluent-bit.conf   -ti cr.fluentbit.io/fluent/fluent-bit:2.0   /fluent-bit/bin/fluent-bit   -c /fluent-bit/etc/fluent-bit.conf