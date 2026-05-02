
/**
 * Samples for Accounts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/DeleteAccount.json
     */
    /**
     * Sample code: Delete Account.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteAccount(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().delete("myResourceGroup", "PropTest01", com.azure.core.util.Context.NONE);
    }
}
