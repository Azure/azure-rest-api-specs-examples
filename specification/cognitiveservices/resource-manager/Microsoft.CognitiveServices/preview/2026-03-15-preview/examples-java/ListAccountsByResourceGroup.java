
/**
 * Samples for Accounts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListAccountsByResourceGroup.json
     */
    /**
     * Sample code: List Accounts by Resource Group.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listAccountsByResourceGroup(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
