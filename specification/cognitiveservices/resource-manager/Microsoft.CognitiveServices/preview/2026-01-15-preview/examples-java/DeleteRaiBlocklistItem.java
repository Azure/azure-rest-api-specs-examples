
/**
 * Samples for RaiBlocklistItems Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/DeleteRaiBlocklistItem.json
     */
    /**
     * Sample code: DeleteRaiBlocklistItem.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteRaiBlocklistItem(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBlocklistItems().delete("resourceGroupName", "accountName", "raiBlocklistName",
            "raiBlocklistItemName", com.azure.core.util.Context.NONE);
    }
}
