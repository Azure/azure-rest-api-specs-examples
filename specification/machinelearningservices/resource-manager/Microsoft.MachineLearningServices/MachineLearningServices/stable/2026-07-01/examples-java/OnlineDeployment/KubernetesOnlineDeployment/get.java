
/**
 * Samples for OnlineDeployments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/OnlineDeployment/KubernetesOnlineDeployment/get.json
     */
    /**
     * Sample code: Get Kubernetes Online Deployment.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getKubernetesOnlineDeployment(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineDeployments().getWithResponse("test-rg", "my-aml-workspace", "testEndpointName",
            "testDeploymentName", com.azure.core.util.Context.NONE);
    }
}
