import com.azure.core.util.Context;

/** Samples for OnlineDeployments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineDeployment/KubernetesOnlineDeployment/get.json
     */
    /**
     * Sample code: Get Kubernetes Online Deployment.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getKubernetesOnlineDeployment(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .onlineDeployments()
            .getWithResponse("test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", Context.NONE);
    }
}
