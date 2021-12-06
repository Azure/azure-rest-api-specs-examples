Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.4/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.videoanalyzer.models.ParameterDefinition;
import java.util.Arrays;

/** Samples for LivePipelines CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/live-pipeline-create.json
     */
    /**
     * Sample code: Create or update a live pipeline.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void createOrUpdateALivePipeline(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .livePipelines()
            .define("livePipeline1")
            .withExistingVideoAnalyzer("testrg", "testaccount2")
            .withTopologyName("pipelinetopology1")
            .withDescription("Live Pipeline 1 Description")
            .withBitrateKbps(500)
            .withParameters(
                Arrays
                    .asList(
                        new ParameterDefinition().withName("rtspUrlParameter").withValue("rtsp://contoso.com/stream")))
            .create();
    }
}
```
