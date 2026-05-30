
/**
 * Samples for RaiBlocklists Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/GetRaiBlocklist.json
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
