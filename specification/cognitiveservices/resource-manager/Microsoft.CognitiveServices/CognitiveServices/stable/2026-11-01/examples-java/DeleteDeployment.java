
/**
 * Samples for Deployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/DeleteDeployment.json
     */
    /**
     * Sample code: DeleteDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void deleteDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deployments().delete("resourceGroupName", "accountName", "deploymentName",
            com.azure.core.util.Context.NONE);
    }
}
