Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.4/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for PipelineTopologies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-get-by-name.json
     */
    /**
     * Sample code: Get a pipeline topology by name.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getAPipelineTopologyByName(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.pipelineTopologies().getWithResponse("testrg", "testaccount2", "pipelineTopology1", Context.NONE);
    }
}
```
