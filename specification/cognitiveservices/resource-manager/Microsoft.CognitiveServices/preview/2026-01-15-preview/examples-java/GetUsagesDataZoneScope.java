
/**
 * Samples for Accounts ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/GetUsagesDataZoneScope.json
     */
    /**
     * Sample code: Get Usages DataZone Scope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getUsagesDataZoneScope(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().listUsagesWithResponse("myResourceGroup", "TestUsage02", null,
            com.azure.core.util.Context.NONE);
    }
}
