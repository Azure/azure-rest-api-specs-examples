import com.azure.core.util.Context;

/** Samples for VideoAnalyzerOperationStatuses Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-operation-status-by-id-non-terminal-state-failed.json
     */
    /**
     * Sample code: Get status of asynchronous operation when it is completed with error.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void getStatusOfAsynchronousOperationWhenItIsCompletedWithError(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager
            .videoAnalyzerOperationStatuses()
            .getWithResponse("westus", "D612C429-2526-49D5-961B-885AE11406FD", Context.NONE);
    }
}
