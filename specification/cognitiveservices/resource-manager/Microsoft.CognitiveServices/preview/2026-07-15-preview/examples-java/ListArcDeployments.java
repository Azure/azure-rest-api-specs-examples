
/**
 * Samples for ArcDeployments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-15-preview/ListArcDeployments.json
     */
    /**
     * Sample code: ListArcDeployments.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listArcDeployments(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.arcDeployments().list("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
