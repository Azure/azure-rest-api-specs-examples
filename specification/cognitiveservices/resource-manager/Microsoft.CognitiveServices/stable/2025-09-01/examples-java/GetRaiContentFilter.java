
/**
 * Samples for RaiContentFilters Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * GetRaiContentFilter.json
     */
    /**
     * Sample code: GetRaiContentFilters.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getRaiContentFilters(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiContentFilters().getWithResponse("WestUS", "IndirectAttack", com.azure.core.util.Context.NONE);
    }
}
