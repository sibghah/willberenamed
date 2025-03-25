import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as ec2 from 'aws-cdk-lib/aws-ec2';
import * as ecs from 'aws-cdk-lib/aws-ecs';
import * as ecr from 'aws-cdk-lib/aws-ecr';
import * as ecs_patterns from 'aws-cdk-lib/aws-ecs-patterns';

export class InfraStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    const vpc = new ec2.Vpc(this, 'Vpc', {
      maxAzs: 3,
    });

    const repository = ecr.Repository.fromRepositoryName(this, 'WorkoutApiRepository', 'workout-api-repository');

    const cluster = new ecs.Cluster(this, 'EcsCluster', {
      vpc,
    });

    const fargateService = new ecs_patterns.ApplicationLoadBalancedFargateService(this, 'WorkoutApiService', {
      cluster,
      cpu: 512,
      memoryLimitMiB: 1024,
      desiredCount: 1,
      taskImageOptions: {
        image: ecs.ContainerImage.fromEcrRepository(repository),
        containerPort: 8080,
      },
      publicLoadBalancer: true,
    });

    // Configure health check path
    fargateService.targetGroup.configureHealthCheck({
      path: '/workouts', // Replace with your actual health check endpoint
      interval: cdk.Duration.seconds(30), // Optional: adjust health check interval
      timeout: cdk.Duration.seconds(5),   // Optional: adjust health check timeout
      healthyThresholdCount: 2,           // Optional: adjust healthy threshold count
      unhealthyThresholdCount: 2,         // Optional: adjust unhealthy threshold count
    });
  }
}
