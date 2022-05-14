Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VideoAnalyzerOperationStatuses Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-operation-status-by-id-terminal-state.json
     */
    /**
     * Sample code: Get status of asynchronous operation when it is completed.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getStatusOfAsynchronousOperationWhenItIsCompleted(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .videoAnalyzerOperationStatuses()
            .getWithResponse("westus", "D612C429-2526-49D5-961B-885AE11406FD", Context.NONE);
    }
}
```
