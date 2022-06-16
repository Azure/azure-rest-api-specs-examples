import com.azure.core.util.Context;

/** Samples for VideoAnalyzers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-accounts-list-all-accounts.json
     */
    /**
     * Sample code: List all Video Analyzer accounts.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void listAllVideoAnalyzerAccounts(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.videoAnalyzers().listWithResponse("contoso", Context.NONE);
    }
}
