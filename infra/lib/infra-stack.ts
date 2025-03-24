import * as cdk from 'aws-cdk-lib';
import { Construct } from 'constructs';
import * as ec2 from 'aws-cdk-lib/aws-ec2';
import * as ecs from 'aws-cdk-lib/aws-ecs';
import * as ecr from 'aws-cdk-lib/aws-ecr';
import * as ecs_patterns from 'aws-cdk-lib/aws-ecs-patterns';

export class InfraStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props?: cdk.StackProps) {
    super(scope, id, props);

    // The code that defines your stack goes here

    // example resource
    // const queue = new sqs.Queue(this, 'InfraQueue', {
    //   visibilityTimeout: cdk.Duration.seconds(300)
    // });

    const vpc = new ec2.Vpc(this, 'Vpc', {
      maxAzs: 3,
    });

    // const repository = new ecr.Repository(this, 'workout-api-repository', {
    //   removalPolicy: cdk.RemovalPolicy.DESTROY, //only for dev/testing
    // });

    const repository = ecr.Repository.fromRepositoryName(this, 'WorkoutApiRepository', 'workout-api-repository');
    

    //the following creates an ECS cluster inside the VPC
    const cluster = new ecs.Cluster(this, 'EcsCluster', {
      vpc,
    });

    new ecs_patterns.ApplicationLoadBalancedFargateService(this, 'WorkoutApiService', {
      cluster,
      cpu: 512,
      memoryLimitMiB: 1024,
      desiredCount: 1, //number of instances of the API
      taskImageOptions: {
        image: ecs.ContainerImage.fromEcrRepository(repository), 
        containerPort: 8080,
      },
      publicLoadBalancer: true, //make the service publicly available
    });
  }
}
