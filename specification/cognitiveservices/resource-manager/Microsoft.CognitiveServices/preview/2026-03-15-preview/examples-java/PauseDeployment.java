
/**
 * Samples for Deployments Pause.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/PauseDeployment.json
     */
    /**
     * Sample code: PauseDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void pauseDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deployments().pauseWithResponse("resourceGroupName", "accountName", "deploymentName",
            com.azure.core.util.Context.NONE);
    }
}
