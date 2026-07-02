
/**
 * Samples for Accounts ListKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ListKeys.json
     */
    /**
     * Sample code: List Keys.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listKeys(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listKeysWithResponse("myResourceGroup", "myAccount", com.azure.core.util.Context.NONE);
    }
}
