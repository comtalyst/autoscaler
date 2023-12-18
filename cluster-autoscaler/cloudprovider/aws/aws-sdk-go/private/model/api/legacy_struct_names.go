package api

var legacyStructNames = map[string]string{
	"accessanalyzer":                  "Access Analyzer",
	"account":                         "AWS Account",
	"acm":                             "ACM",
	"acm pca":                         "ACM-PCA",
	"alexa for business":              "Alexa For Business",
	"amp":                             "Amazon Prometheus Service",
	"amplify":                         "Amplify",
	"amplifybackend":                  "AmplifyBackend",
	"amplifyuibuilder":                "AWS Amplify UI Builder",
	"api gateway":                     "Amazon API Gateway",
	"apigatewaymanagementapi":         "AmazonApiGatewayManagementApi",
	"apigatewayv2":                    "AmazonApiGatewayV2",
	"app mesh":                        "AWS App Mesh",
	"appconfig":                       "AppConfig",
	"appconfigdata":                   "AWS AppConfig Data",
	"appflow":                         "Amazon Appflow",
	"appintegrations":                 "Amazon AppIntegrations Service",
	"application auto scaling":        "Application Auto Scaling",
	"application discovery service":   "AWS Application Discovery Service",
	"application insights":            "Application Insights",
	"applicationcostprofiler":         "AWS Application Cost Profiler",
	"apprunner":                       "AWS App Runner",
	"appstream":                       "Amazon AppStream",
	"appsync":                         "AWSAppSync",
	"athena":                          "Amazon Athena",
	"auditmanager":                    "AWS Audit Manager",
	"auto scaling":                    "Auto Scaling",
	"auto scaling plans":              "AWS Auto Scaling Plans",
	"backup":                          "AWS Backup",
	"backup gateway":                  "AWS Backup Gateway",
	"batch":                           "AWS Batch",
	"billingconductor":                "AWSBillingConductor",
	"braket":                          "Braket",
	"budgets":                         "AWSBudgets",
	"chime":                           "Amazon Chime",
	"chime sdk identity":              "Amazon Chime SDK Identity",
	"chime sdk media pipelines":       "Amazon Chime SDK Media Pipelines",
	"chime sdk meetings":              "Amazon Chime SDK Meetings",
	"chime sdk messaging":             "Amazon Chime SDK Messaging",
	"cloud9":                          "AWS Cloud9",
	"cloudcontrol":                    "CloudControlApi",
	"clouddirectory":                  "Amazon CloudDirectory",
	"cloudformation":                  "AWS CloudFormation",
	"cloudfront":                      "CloudFront",
	"cloudhsm":                        "CloudHSM",
	"cloudhsm v2":                     "CloudHSM V2",
	"cloudsearch":                     "Amazon CloudSearch",
	"cloudsearch domain":              "Amazon CloudSearch Domain",
	"cloudtrail":                      "CloudTrail",
	"cloudwatch":                      "CloudWatch",
	"cloudwatch events":               "Amazon CloudWatch Events",
	"cloudwatch logs":                 "Amazon CloudWatch Logs",
	"codeartifact":                    "CodeArtifact",
	"codebuild":                       "AWS CodeBuild",
	"codecommit":                      "CodeCommit",
	"codedeploy":                      "CodeDeploy",
	"codeguru reviewer":               "CodeGuruReviewer",
	"codeguruprofiler":                "Amazon CodeGuru Profiler",
	"codepipeline":                    "CodePipeline",
	"codestar":                        "CodeStar",
	"codestar connections":            "AWS CodeStar connections",
	"codestar notifications":          "AWS CodeStar Notifications",
	"cognito identity":                "Amazon Cognito Identity",
	"cognito identity provider":       "Amazon Cognito Identity Provider",
	"cognito sync":                    "Amazon Cognito Sync",
	"comprehend":                      "Amazon Comprehend",
	"comprehendmedical":               "ComprehendMedical",
	"compute optimizer":               "AWS Compute Optimizer",
	"config service":                  "Config Service",
	"connect":                         "Amazon Connect",
	"connect contact lens":            "Amazon Connect Contact Lens",
	"connectparticipant":              "Amazon Connect Participant",
	"cost and usage report service":   "AWS Cost and Usage Report Service",
	"cost explorer":                   "AWS Cost Explorer",
	"customer profiles":               "Customer Profiles",
	"data pipeline":                   "AWS Data Pipeline",
	"database migration service":      "AWS Database Migration Service",
	"databrew":                        "AWS Glue DataBrew",
	"dataexchange":                    "AWS Data Exchange",
	"datasync":                        "DataSync",
	"dax":                             "Amazon DAX",
	"detective":                       "Amazon Detective",
	"device farm":                     "AWS Device Farm",
	"devops guru":                     "Amazon DevOps Guru",
	"direct connect":                  "AWS Direct Connect",
	"directory service":               "Directory Service",
	"dlm":                             "Amazon DLM",
	"docdb":                           "Amazon DocDB",
	"drs":                             "drs",
	"dynamodb":                        "DynamoDB",
	"dynamodb streams":                "Amazon DynamoDB Streams",
	"ebs":                             "Amazon EBS",
	"ec2":                             "Amazon EC2",
	"ec2 instance connect":            "EC2 Instance Connect",
	"ecr":                             "Amazon ECR",
	"ecr public":                      "Amazon ECR Public",
	"ecs":                             "Amazon ECS",
	"efs":                             "EFS",
	"eks":                             "Amazon EKS",
	"elastic beanstalk":               "Elastic Beanstalk",
	"elastic inference":               "Amazon Elastic Inference",
	"elastic load balancing":          "Elastic Load Balancing",
	"elastic load balancing v2":       "Elastic Load Balancing v2",
	"elastic transcoder":              "Amazon Elastic Transcoder",
	"elasticache":                     "Amazon ElastiCache",
	"elasticsearch service":           "Amazon Elasticsearch Service",
	"emr":                             "Amazon EMR",
	"emr containers":                  "Amazon EMR Containers",
	"emr serverless":                  "EMR Serverless",
	"eventbridge":                     "Amazon EventBridge",
	"evidently":                       "Amazon CloudWatch Evidently",
	"finspace":                        "finspace",
	"finspace data":                   "FinSpace Data",
	"firehose":                        "Firehose",
	"fis":                             "FIS",
	"fms":                             "FMS",
	"forecast":                        "Amazon Forecast Service",
	"forecastquery":                   "Amazon Forecast Query Service",
	"frauddetector":                   "Amazon Fraud Detector",
	"fsx":                             "Amazon FSx",
	"gamelift":                        "Amazon GameLift",
	"gamesparks":                      "GameSparks",
	"glacier":                         "Amazon Glacier",
	"global accelerator":              "AWS Global Accelerator",
	"glue":                            "AWS Glue",
	"grafana":                         "Amazon Managed Grafana",
	"greengrass":                      "AWS Greengrass",
	"greengrassv2":                    "AWS GreengrassV2",
	"groundstation":                   "AWS Ground Station",
	"guardduty":                       "Amazon GuardDuty",
	"health":                          "AWSHealth",
	"healthlake":                      "HealthLake",
	"honeycode":                       "Honeycode",
	"iam":                             "IAM",
	"identitystore":                   "IdentityStore",
	"imagebuilder":                    "imagebuilder",
	"import export":                   "AWS Import/Export",
	"inspector":                       "Amazon Inspector",
	"inspector2":                      "Inspector2",
	"iot":                             "AWS IoT",
	"iot 1click devices service":      "AWS IoT 1-Click Devices Service",
	"iot 1click projects":             "AWS IoT 1-Click Projects",
	"iot data plane":                  "AWS IoT Data Plane",
	"iot events":                      "AWS IoT Events",
	"iot events data":                 "AWS IoT Events Data",
	"iot jobs data plane":             "AWS IoT Jobs Data Plane",
	"iot wireless":                    "AWS IoT Wireless",
	"iotanalytics":                    "AWS IoT Analytics",
	"iotdeviceadvisor":                "AWSIoTDeviceAdvisor",
	"iotfleethub":                     "AWS IoT Fleet Hub",
	"iotsecuretunneling":              "AWS IoT Secure Tunneling",
	"iotsitewise":                     "AWS IoT SiteWise",
	"iotthingsgraph":                  "AWS IoT Things Graph",
	"iottwinmaker":                    "AWS IoT TwinMaker",
	"ivs":                             "Amazon IVS",
	"ivschat":                         "ivschat",
	"kafka":                           "Kafka",
	"kafkaconnect":                    "Kafka Connect",
	"kendra":                          "kendra",
	"keyspaces":                       "Amazon Keyspaces",
	"kinesis":                         "Kinesis",
	"kinesis analytics":               "Kinesis Analytics",
	"kinesis analytics v2":            "Kinesis Analytics V2",
	"kinesis video":                   "Kinesis Video",
	"kinesis video archived media":    "Kinesis Video Archived Media",
	"kinesis video media":             "Kinesis Video Media",
	"kinesis video signaling":         "Amazon Kinesis Video Signaling Channels",
	"kms":                             "KMS",
	"lakeformation":                   "AWS Lake Formation",
	"lambda":                          "AWS Lambda",
	"lex model building service":      "Amazon Lex Model Building Service",
	"lex models v2":                   "Lex Models V2",
	"lex runtime service":             "Amazon Lex Runtime Service",
	"lex runtime v2":                  "Lex Runtime V2",
	"license manager":                 "AWS License Manager",
	"lightsail":                       "Amazon Lightsail",
	"location":                        "Amazon Location Service",
	"lookoutequipment":                "LookoutEquipment",
	"lookoutmetrics":                  "LookoutMetrics",
	"lookoutvision":                   "Amazon Lookout for Vision",
	"machine learning":                "Amazon Machine Learning",
	"macie":                           "Amazon Macie",
	"macie2":                          "Amazon Macie 2",
	"managedblockchain":               "ManagedBlockchain",
	"marketplace catalog":             "AWS Marketplace Catalog",
	"marketplace commerce analytics":  "AWS Marketplace Commerce Analytics",
	"marketplace entitlement service": "AWS Marketplace Entitlement Service",
	"marketplace metering":            "AWSMarketplace Metering",
	"mediaconnect":                    "AWS MediaConnect",
	"mediaconvert":                    "MediaConvert",
	"medialive":                       "MediaLive",
	"mediapackage":                    "MediaPackage",
	"mediapackage vod":                "MediaPackage Vod",
	"mediastore":                      "MediaStore",
	"mediastore data":                 "MediaStore Data",
	"mediatailor":                     "MediaTailor",
	"memorydb":                        "Amazon MemoryDB",
	"mgn":                             "mgn",
	"migration hub":                   "AWS Migration Hub",
	"migration hub refactor spaces":   "AWS Migration Hub Refactor Spaces",
	"migrationhub config":             "AWS Migration Hub Config",
	"migrationhubstrategy":            "Migration Hub Strategy Recommendations",
	"mobile":                          "AWS Mobile",
	"mobile analytics":                "Amazon Mobile Analytics",
	"mq":                              "AmazonMQ",
	"mturk":                           "Amazon MTurk",
	"mwaa":                            "AmazonMWAA",
	"neptune":                         "Amazon Neptune",
	"network firewall":                "Network Firewall",
	"networkmanager":                  "NetworkManager",
	"nimble":                          "AmazonNimbleStudio",
	"opensearch":                      "Amazon OpenSearch Service",
	"opsworks":                        "AWS OpsWorks",
	"opsworkscm":                      "OpsWorksCM",
	"organizations":                   "Organizations",
	"outposts":                        "Outposts",
	"panorama":                        "Panorama",
	"personalize":                     "Amazon Personalize",
	"personalize events":              "Amazon Personalize Events",
	"personalize runtime":             "Amazon Personalize Runtime",
	"pi":                              "AWS PI",
	"pinpoint":                        "Amazon Pinpoint",
	"pinpoint email":                  "Pinpoint Email",
	"pinpoint sms voice":              "Pinpoint SMS Voice",
	"pinpoint sms voice v2":           "Amazon Pinpoint SMS Voice V2",
	"polly":                           "Amazon Polly",
	"pricing":                         "AWS Pricing",
	"proton":                          "AWS Proton",
	"qldb":                            "QLDB",
	"qldb session":                    "QLDB Session",
	"quicksight":                      "Amazon QuickSight",
	"ram":                             "RAM",
	"rbin":                            "Amazon Recycle Bin",
	"rds":                             "Amazon RDS",
	"rds data":                        "AWS RDS DataService",
	"redshift":                        "Amazon Redshift",
	"redshift data":                   "Redshift Data API Service",
	"rekognition":                     "Amazon Rekognition",
	"resiliencehub":                   "AWS Resilience Hub",
	"resource groups":                 "Resource Groups",
	"resource groups tagging api":     "AWS Resource Groups Tagging API",
	"robomaker":                       "RoboMaker",
	"route 53":                        "Route 53",
	"route 53 domains":                "Amazon Route 53 Domains",
	"route53 recovery cluster":        "Route53 Recovery Cluster",
	"route53 recovery control config": "AWS Route53 Recovery Control Config",
	"route53 recovery readiness":      "AWS Route53 Recovery Readiness",
	"route53resolver":                 "Route53Resolver",
	"rum":                             "CloudWatch RUM",
	"s3":                              "Amazon S3",
	"s3 control":                      "AWS S3 Control",
	"s3outposts":                      "Amazon S3 Outposts",
	"sagemaker":                       "SageMaker",
	"sagemaker a2i runtime":           "Amazon Augmented AI Runtime",
	"sagemaker edge":                  "Amazon Sagemaker Edge Manager",
	"sagemaker featurestore runtime":  "Amazon SageMaker Feature Store Runtime",
	"sagemaker runtime":               "Amazon SageMaker Runtime",
	"savingsplans":                    "AWSSavingsPlans",
	"schemas":                         "Schemas",
	"secrets manager":                 "AWS Secrets Manager",
	"securityhub":                     "AWS SecurityHub",
	"serverlessapplicationrepository": "AWSServerlessApplicationRepository",
	"service catalog":                 "AWS Service Catalog",
	"service catalog appregistry":     "AppRegistry",
	"service quotas":                  "Service Quotas",
	"servicediscovery":                "ServiceDiscovery",
	"ses":                             "Amazon SES",
	"sesv2":                           "Amazon SES V2",
	"sfn":                             "AWS SFN",
	"shield":                          "AWS Shield",
	"signer":                          "signer",
	"simpledb":                        "Amazon SimpleDB",
	"sms":                             "SMS",
	"snow device management":          "AWS Snow Device Management",
	"snowball":                        "Amazon Snowball",
	"sns":                             "Amazon SNS",
	"sqs":                             "Amazon SQS",
	"ssm":                             "Amazon SSM",
	"ssm contacts":                    "SSM Contacts",
	"ssm incidents":                   "SSM Incidents",
	"sso":                             "SSO",
	"sso admin":                       "SSO Admin",
	"sso oidc":                        "SSO OIDC",
	"storage gateway":                 "AWS Storage Gateway",
	"sts":                             "AWS STS",
	"support":                         "AWS Support",
	"swf":                             "Amazon SWF",
	"synthetics":                      "Synthetics",
	"textract":                        "Amazon Textract",
	"timestream query":                "Timestream Query",
	"timestream write":                "Timestream Write",
	"transcribe":                      "Amazon Transcribe Service",
	"transcribe streaming":            "Amazon Transcribe Streaming Service",
	"transfer":                        "AWS Transfer",
	"translate":                       "Amazon Translate",
	"voice id":                        "Amazon Voice ID",
	"waf":                             "WAF",
	"waf regional":                    "WAF Regional",
	"wafv2":                           "WAFV2",
	"wellarchitected":                 "Well-Architected",
	"wisdom":                          "Amazon Connect Wisdom Service",
	"workdocs":                        "Amazon WorkDocs",
	"worklink":                        "WorkLink",
	"workmail":                        "Amazon WorkMail",
	"workmailmessageflow":             "Amazon WorkMail Message Flow",
	"workspaces":                      "Amazon WorkSpaces",
	"workspaces web":                  "Amazon WorkSpaces Web",
	"xray":                            "AWS X-Ray",
}
