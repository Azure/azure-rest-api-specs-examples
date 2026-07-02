
/**
 * Samples for Accounts ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/GetUsages.json
     */
    /**
     * Sample code: Get Usages.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getUsages(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listUsagesWithResponse("myResourceGroup", "TestUsage02", null,
            com.azure.core.util.Context.NONE);
    }
}
