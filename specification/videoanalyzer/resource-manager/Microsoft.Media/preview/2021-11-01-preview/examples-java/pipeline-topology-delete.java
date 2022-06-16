import com.azure.core.util.Context;

/** Samples for PipelineTopologies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-delete.json
     */
    /**
     * Sample code: Delete a pipeline topology.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void deleteAPipelineTopology(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.pipelineTopologies().deleteWithResponse("testrg", "testaccount2", "pipelineTopology1", Context.NONE);
    }
}
