Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.4/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VideoAnalyzerOperationResults Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-operation-result-by-id.json
     */
    /**
     * Sample code: Get status of asynchronous operation.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getStatusOfAsynchronousOperation(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .videoAnalyzerOperationResults()
            .getWithResponse("westus", "6FBA62C4-99B5-4FF8-9826-FC4744A8864F", Context.NONE);
    }
}
```
