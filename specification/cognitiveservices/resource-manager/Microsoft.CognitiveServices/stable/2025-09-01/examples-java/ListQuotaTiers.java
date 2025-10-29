
/**
 * Samples for QuotaTiers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * ListQuotaTiers.json
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
