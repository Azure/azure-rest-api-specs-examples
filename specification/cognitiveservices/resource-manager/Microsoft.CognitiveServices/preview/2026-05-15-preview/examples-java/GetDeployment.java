
/**
 * Samples for Deployments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/GetDeployment.json
     */
    /**
     * Sample code: GetDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void getDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deployments().getWithResponse("resourceGroupName", "accountName", "deploymentName",
            com.azure.core.util.Context.NONE);
    }
}
