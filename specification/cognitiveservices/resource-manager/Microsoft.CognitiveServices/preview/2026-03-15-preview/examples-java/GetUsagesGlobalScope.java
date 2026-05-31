
/**
 * Samples for Accounts ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/GetUsagesGlobalScope.json
     */
    /**
     * Sample code: Get Usages Global Scope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getUsagesGlobalScope(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listUsagesWithResponse("myResourceGroup", "TestUsage02", null,
            com.azure.core.util.Context.NONE);
    }
}
