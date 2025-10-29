
/**
 * Samples for Deployments ListSkus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * ListDeploymentSkus.json
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
