
/**
 * Samples for QuotaTiers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * GetQuotaTier.json
     */
    /**
     * Sample code: Get the Quota Tier information for a subscription.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getTheQuotaTierInformationForASubscription(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.quotaTiers().getWithResponse("default", com.azure.core.util.Context.NONE);
    }
}
