
/**
 * Samples for RaiToolLabels List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/ListRaiToolLabels.json
     */
    /**
     * Sample code: ListRaiToolLabels.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listRaiToolLabels(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.raiToolLabels().list("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
