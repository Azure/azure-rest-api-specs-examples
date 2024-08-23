
/**
 * Samples for OnlineDeployments ListSkus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/OnlineDeployment/ManagedOnlineDeployment/listSkus.json
     */
    /**
     * Sample code: List Managed Online Deployment Skus.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listManagedOnlineDeploymentSkus(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineDeployments().listSkus("test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", 1,
            null, com.azure.core.util.Context.NONE);
    }
}
