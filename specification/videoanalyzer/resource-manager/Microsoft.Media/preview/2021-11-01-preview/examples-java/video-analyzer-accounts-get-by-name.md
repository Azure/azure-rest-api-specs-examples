```java
import com.azure.core.util.Context;

/** Samples for VideoAnalyzers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-accounts-get-by-name.json
     */
    /**
     * Sample code: Get a Video Analyzer account by name.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getAVideoAnalyzerAccountByName(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.videoAnalyzers().getByResourceGroupWithResponse("contoso", "contosotv", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.
