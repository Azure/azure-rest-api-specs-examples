
/**
 * Samples for Deployments ListSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ListDeploymentSkus.json
     */
    /**
     * Sample code: ListDeploymentSkus.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listDeploymentSkus(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deployments().listSkus("resourceGroupName", "accountName", "deploymentName",
            com.azure.core.util.Context.NONE);
    }
}
