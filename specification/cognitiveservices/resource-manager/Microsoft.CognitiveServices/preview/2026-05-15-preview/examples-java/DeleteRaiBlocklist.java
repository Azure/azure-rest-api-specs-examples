
/**
 * Samples for RaiBlocklists Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/DeleteRaiBlocklist.json
     */
    /**
     * Sample code: DeleteRaiBlocklist.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteRaiBlocklist(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiBlocklists().delete("resourceGroupName", "accountName", "raiBlocklistName",
            com.azure.core.util.Context.NONE);
    }
}
