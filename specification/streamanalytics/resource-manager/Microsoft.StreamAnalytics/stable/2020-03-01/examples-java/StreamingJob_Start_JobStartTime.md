```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.streamanalytics.models.OutputStartMode;
import com.azure.resourcemanager.streamanalytics.models.StartStreamingJobParameters;

/** Samples for StreamingJobs Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/StreamingJob_Start_JobStartTime.json
     */
    /**
     * Sample code: Start a streaming job with JobStartTime output start mode.
     *
     * @param manager Entry point to StreamAnalyticsManager.
     */
    public static void startAStreamingJobWithJobStartTimeOutputStartMode(
        com.azure.resourcemanager.streamanalytics.StreamAnalyticsManager manager) {
        manager
            .streamingJobs()
            .start(
                "sjrg6936",
                "sj59",
                new StartStreamingJobParameters().withOutputStartMode(OutputStartMode.JOB_START_TIME),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-streamanalytics_1.0.0-beta.2/sdk/streamanalytics/azure-resourcemanager-streamanalytics/README.md) on how to add the SDK to your project and authenticate.
