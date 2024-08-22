
/**
 * Samples for OnlineDeployments Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/OnlineDeployment/ManagedOnlineDeployment/get.json
     */
    /**
     * Sample code: Get Managed Online Deployment.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getManagedOnlineDeployment(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineDeployments().getWithResponse("test-rg", "my-aml-workspace", "testEndpointName",
            "testDeploymentName", com.azure.core.util.Context.NONE);
    }
}
