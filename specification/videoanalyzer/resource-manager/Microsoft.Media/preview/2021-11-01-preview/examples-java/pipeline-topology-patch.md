```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.videoanalyzer.models.PipelineTopology;

/** Samples for PipelineTopologies Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-patch.json
     */
    /**
     * Sample code: Update pipeline topology.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void updatePipelineTopology(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        PipelineTopology resource =
            manager
                .pipelineTopologies()
                .getWithResponse("testrg", "testaccount2", "pipelineTopology1", Context.NONE)
                .getValue();
        resource.update().withDescription("Pipeline Topology 1 Description").apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-videoanalyzer_1.0.0-beta.5/sdk/videoanalyzer/azure-resourcemanager-videoanalyzer/README.md) on how to add the SDK to your project and authenticate.
