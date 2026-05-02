
/**
 * Samples for Deployments Resume.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/ResumeDeployment.json
     */
    /**
     * Sample code: ResumeDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void resumeDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deployments().resumeWithResponse("resourceGroupName", "accountName", "deploymentName",
            com.azure.core.util.Context.NONE);
    }
}
