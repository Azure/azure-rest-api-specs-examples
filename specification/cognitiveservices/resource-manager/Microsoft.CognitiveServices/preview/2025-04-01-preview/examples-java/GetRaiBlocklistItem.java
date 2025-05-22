
/**
 * Samples for RaiBlocklistItems Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
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
