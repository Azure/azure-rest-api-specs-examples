Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.4/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for VideoAnalyzers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-accounts-delete.json
     */
    /**
     * Sample code: Delete a Video Analyzer account.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void deleteAVideoAnalyzerAccount(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.videoAnalyzers().deleteWithResponse("contoso", "contosotv", Context.NONE);
    }
}
```
