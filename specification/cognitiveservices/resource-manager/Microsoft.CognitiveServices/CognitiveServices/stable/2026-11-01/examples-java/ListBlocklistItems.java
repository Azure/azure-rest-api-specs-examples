
/**
 * Samples for RaiBlocklistItems List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ListBlocklistItems.json
     */
    /**
     * Sample code: ListBlocklistItems.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listBlocklistItems(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBlocklistItems().list("resourceGroupName", "accountName", "raiBlocklistName",
            com.azure.core.util.Context.NONE);
    }
}
