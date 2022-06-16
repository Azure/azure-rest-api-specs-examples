import com.azure.core.util.Context;

/** Samples for LivePipelines Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/live-pipeline-get-by-name.json
     */
    /**
     * Sample code: Retrieves a specific live pipeline by name.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void retrievesASpecificLivePipelineByName(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.livePipelines().getWithResponse("testrg", "testaccount2", "livePipeline1", Context.NONE);
    }
}
