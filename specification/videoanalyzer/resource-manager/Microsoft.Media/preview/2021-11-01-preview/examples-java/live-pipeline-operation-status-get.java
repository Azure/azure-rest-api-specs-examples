import com.azure.core.util.Context;

/** Samples for LivePipelineOperationStatuses Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/live-pipeline-operation-status-get.json
     */
    /**
     * Sample code: Get the live pipeline operation status.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getTheLivePipelineOperationStatus(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .livePipelineOperationStatuses()
            .getWithResponse(
                "testrg", "testaccount2", "livePipeline1", "00000000-0000-0000-0000-000000000001", Context.NONE);
    }
}
