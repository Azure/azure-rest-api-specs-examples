
/**
 * Samples for Accounts ListModels.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ListAccountModels.json
     */
    /**
     * Sample code: List AccountModels.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listAccountModels(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listModels("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
