import com.azure.core.util.Context;
import com.azure.resourcemanager.videoanalyzer.models.LivePipeline;

/** Samples for LivePipelines Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/live-pipeline-patch.json
     */
    /**
     * Sample code: Updates a live pipeline.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void updatesALivePipeline(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        LivePipeline resource =
            manager.livePipelines().getWithResponse("testrg", "testaccount2", "livePipeline1", Context.NONE).getValue();
        resource.update().withDescription("Live Pipeline 1 Description").apply();
    }
}
