
/**
 * Samples for RaiBlocklistItems Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * GetRaiBlocklistItem.json
     */
    /**
     * Sample code: GetRaiBlocklistItem.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getRaiBlocklistItem(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBlocklistItems().getWithResponse("resourceGroupName", "accountName", "raiBlocklistName",
            "raiBlocklistItemName", com.azure.core.util.Context.NONE);
    }
}
