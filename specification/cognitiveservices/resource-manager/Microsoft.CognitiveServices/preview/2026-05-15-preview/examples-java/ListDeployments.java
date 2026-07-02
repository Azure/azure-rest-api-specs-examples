
/**
 * Samples for Deployments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ListDeployments.json
     */
    /**
     * Sample code: ListDeployments.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listDeployments(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deployments().list("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
