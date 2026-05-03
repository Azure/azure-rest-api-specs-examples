
/**
 * Samples for Accounts ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/GetUsagesClassicScope.json
     */
    /**
     * Sample code: Get Usages Classic Scope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getUsagesClassicScope(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listUsagesWithResponse("myResourceGroup", "TestUsage02", null,
            com.azure.core.util.Context.NONE);
    }
}
