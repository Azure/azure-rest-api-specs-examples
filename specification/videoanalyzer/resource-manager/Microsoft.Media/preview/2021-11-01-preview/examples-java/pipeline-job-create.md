```java
import com.azure.resourcemanager.videoanalyzer.models.ParameterDefinition;
import java.util.Arrays;

/** Samples for PipelineJobs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-job-create.json
     */
    /**
     * Sample code: Create or update a pipeline job.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void createOrUpdateAPipelineJob(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .pipelineJobs()
            .define("pipelineJob1")
            .withExistingVideoAnalyzer("testrg", "testaccount2")
            .withTopologyName("pipelinetopology1")
            .withDescription("Pipeline Job 1 Dsecription")
            .withParameters(
                Arrays
                    .asList(
                        new ParameterDefinition()
                            .withName("timesequences")
                            .withValue("[[\"2020-10-05T03:30:00Z\", \"2020-10-05T04:30:00Z\"]]"),
                        new ParameterDefinition().withName("videoSourceName").withValue("camera001")))
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.
