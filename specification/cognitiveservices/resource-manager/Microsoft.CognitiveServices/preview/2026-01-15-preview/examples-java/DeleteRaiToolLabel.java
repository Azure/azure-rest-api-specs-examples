
/**
 * Samples for RaiToolLabels Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/DeleteRaiToolLabel.json
     */
    /**
     * Sample code: DeleteRaiToolLabel.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteRaiToolLabel(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiToolLabels().delete("resourceGroupName", "accountName", "Web_Search",
            com.azure.core.util.Context.NONE);
    }
}
