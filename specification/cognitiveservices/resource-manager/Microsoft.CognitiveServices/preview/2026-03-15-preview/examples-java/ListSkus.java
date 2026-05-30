
/**
 * Samples for Accounts ListSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListSkus.json
     */
    /**
     * Sample code: List SKUs.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listSKUs(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listSkusWithResponse("myResourceGroup", "myAccount", com.azure.core.util.Context.NONE);
    }
}
