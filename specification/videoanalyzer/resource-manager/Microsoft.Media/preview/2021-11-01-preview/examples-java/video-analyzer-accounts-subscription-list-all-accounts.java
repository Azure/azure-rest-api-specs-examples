import com.azure.core.util.Context;

/** Samples for VideoAnalyzers ListBySubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/videoanalyzer/resource-manager/Microsoft.Media/preview/2021-11-01-preview/examples/video-analyzer-accounts-subscription-list-all-accounts.json
     */
    /**
     * Sample code: List all Video Analyzer accounts in the specified subscription.
     *
     * @param manager Entry point to VideoAnalyzerManager.
     */
    public static void listAllVideoAnalyzerAccountsInTheSpecifiedSubscription(
        com.azure.resourcemanager.videoanalyzer.VideoAnalyzerManager manager) {
        manager.videoAnalyzers().listBySubscriptionWithResponse(Context.NONE);
    }
}
