
/**
 * Samples for RaiContentFilters List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * ListRaiContentFilters.json
     */
    /**
     * Sample code: ListRaiContentFilters.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listRaiContentFilters(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiContentFilters().list("WestUS", com.azure.core.util.Context.NONE);
    }
}
