
/**
 * Samples for RaiBlocklistItems Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/GetRaiBlocklistItem.json
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
