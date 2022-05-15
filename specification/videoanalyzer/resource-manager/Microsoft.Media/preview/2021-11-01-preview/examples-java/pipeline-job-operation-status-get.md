Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PipelineJobOperationStatuses Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-job-operation-status-get.json
     */
    /**
     * Sample code: Get the pipeline job operation status.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getThePipelineJobOperationStatus(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .pipelineJobOperationStatuses()
            .getWithResponse(
                "testrg", "testaccount2", "pipelineJob1", "00000000-0000-0000-0000-000000000001", Context.NONE);
    }
}
```
