
/**
 * Samples for RaiBlocklists Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * GetRaiBlocklist.json
     */
    /**
     * Sample code: GetRaiBlocklist.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getRaiBlocklist(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBlocklists().getWithResponse("resourceGroupName", "accountName", "raiBlocklistName",
            com.azure.core.util.Context.NONE);
    }
}
