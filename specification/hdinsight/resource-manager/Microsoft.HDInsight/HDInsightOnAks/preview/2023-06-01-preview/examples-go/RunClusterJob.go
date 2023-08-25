package armhdinsightcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hdinsightcontainers/armhdinsightcontainers"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7f70e351393addbc31d790a908c994c7c8644d9c/specification/hdinsight/resource-manager/Microsoft.HDInsight/HDInsightOnAks/preview/2023-06-01-preview/examples/RunClusterJob.json
func ExampleClusterJobsClient_BeginRunJob() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhdinsightcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewClusterJobsClient().BeginRunJob(ctx, "hiloResourcegroup", "clusterpool1", "cluster1", armhdinsightcontainers.ClusterJob{
		Properties: &armhdinsightcontainers.FlinkJobProperties{
			JobType:    to.Ptr(armhdinsightcontainers.JobTypeFlinkJob),
			Action:     to.Ptr(armhdinsightcontainers.ActionSTART),
			EntryClass: to.Ptr("com.microsoft.hilo.flink.job.streaming.SleepJob"),
			FlinkConfiguration: map[string]*string{
				"parallelism":         to.Ptr("1"),
				"savepoint.directory": to.Ptr("abfs://flinkjob@hilosa.dfs.core.windows.net/savepoint"),
			},
			JarName:         to.Ptr("flink-sleep-job-0.0.1-SNAPSHOT.jar"),
			JobJarDirectory: to.Ptr("abfs://flinkjob@hilosa.dfs.core.windows.net/jars"),
			JobName:         to.Ptr("flink-job-name"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ClusterJob = armhdinsightcontainers.ClusterJob{
	// 	Properties: &armhdinsightcontainers.FlinkJobProperties{
	// 		JobType: to.Ptr(armhdinsightcontainers.JobTypeFlinkJob),
	// 		Action: to.Ptr(armhdinsightcontainers.ActionSTART),
	// 		EntryClass: to.Ptr("com.microsoft.hilo.flink.job.streaming.SleepJob"),
	// 		FlinkConfiguration: map[string]*string{
	// 			"parallelism": to.Ptr("1"),
	// 			"savepoint.directory": to.Ptr("abfs://flinkjob@hilosa.dfs.core.windows.net/savepoint"),
	// 		},
	// 		JarName: to.Ptr("flink-sleep-job-0.0.1-SNAPSHOT.jar"),
	// 		JobJarDirectory: to.Ptr("abfs://flinkjob@hilosa.dfs.core.windows.net/jars"),
	// 		JobName: to.Ptr("flink-job-name"),
	// 	},
	// }
}
