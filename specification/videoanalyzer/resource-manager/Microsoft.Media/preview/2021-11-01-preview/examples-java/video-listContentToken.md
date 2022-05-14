Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Videos ListContentToken. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-listContentToken.json
     */
    /**
     * Sample code: Generate a content token for media endpoint authorization.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void generateAContentTokenForMediaEndpointAuthorization(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.videos().listContentTokenWithResponse("testrg", "testaccount2", "video3", Context.NONE);
    }
}
```
