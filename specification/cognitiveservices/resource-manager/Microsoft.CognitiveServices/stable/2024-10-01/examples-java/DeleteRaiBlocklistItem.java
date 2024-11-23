
/**
 * Samples for RaiBlocklistItems Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/
     * DeleteRaiBlocklistItem.json
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
