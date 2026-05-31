
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListUsagesDataZoneScope.json
     */
    /**
     * Sample code: Get Usages DataZone Scope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getUsagesDataZoneScope(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.usages().list("WestUS", null, com.azure.core.util.Context.NONE);
    }
}
