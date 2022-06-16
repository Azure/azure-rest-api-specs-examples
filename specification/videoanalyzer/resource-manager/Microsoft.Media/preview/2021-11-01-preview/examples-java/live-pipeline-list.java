import com.azure.core.util.Context;

/** Samples for LivePipelines List. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/live-pipeline-list.json
     */
    /**
     * Sample code: List live pipelines.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void listLivePipelines(com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.livePipelines().list("testrg", "testaccount2", null, 2, Context.NONE);
    }
}
