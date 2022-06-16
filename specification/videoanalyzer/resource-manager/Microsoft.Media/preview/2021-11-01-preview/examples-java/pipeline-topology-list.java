import com.azure.core.util.Context;

/** Samples for PipelineTopologies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/pipeline-topology-list.json
     */
    /**
     * Sample code: List all pipeline topologies.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void listAllPipelineTopologies(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.pipelineTopologies().list("testrg", "testaccount2", null, 2, Context.NONE);
    }
}
