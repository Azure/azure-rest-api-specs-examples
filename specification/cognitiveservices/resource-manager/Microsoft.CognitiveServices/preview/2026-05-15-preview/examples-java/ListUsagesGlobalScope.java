
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ListUsagesGlobalScope.json
     */
    /**
     * Sample code: Get Usages Global Scope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getUsagesGlobalScope(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.usages().list("WestUS", null, com.azure.core.util.Context.NONE);
    }
}
