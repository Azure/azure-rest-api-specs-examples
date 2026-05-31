
/**
 * Samples for QuotaTiers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListQuotaTiers.json
     */
    /**
     * Sample code: List the Quota Tier for a subscription.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listTheQuotaTierForASubscription(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.quotaTiers().list(com.azure.core.util.Context.NONE);
    }
}
