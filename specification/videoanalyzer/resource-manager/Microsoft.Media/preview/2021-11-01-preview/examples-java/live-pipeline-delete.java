import com.azure.core.util.Context;

/** Samples for LivePipelines Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/live-pipeline-delete.json
     */
    /**
     * Sample code: Delete a live pipeline.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void deleteALivePipeline(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.livePipelines().deleteWithResponse("testrg", "testaccount2", "livePipeline1", Context.NONE);
    }
}
