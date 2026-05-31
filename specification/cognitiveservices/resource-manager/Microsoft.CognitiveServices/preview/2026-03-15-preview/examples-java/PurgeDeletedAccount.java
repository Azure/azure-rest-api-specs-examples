
/**
 * Samples for DeletedAccounts Purge.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/PurgeDeletedAccount.json
     */
    /**
     * Sample code: Delete Account.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteAccount(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deletedAccounts().purge("westus", "myResourceGroup", "PropTest01", com.azure.core.util.Context.NONE);
    }
}
