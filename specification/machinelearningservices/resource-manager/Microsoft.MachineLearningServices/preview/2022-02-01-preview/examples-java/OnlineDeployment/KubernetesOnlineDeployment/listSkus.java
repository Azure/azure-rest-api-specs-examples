import com.azure.core.util.Context;

/** Samples for OnlineDeployments ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineDeployment/KubernetesOnlineDeployment/listSkus.json
     */
    /**
     * Sample code: List Kubernetes Online Deployment Skus.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listKubernetesOnlineDeploymentSkus(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .onlineDeployments()
            .listSkus("test-rg", "my-aml-workspace", "testEndpointName", "testDeploymentName", 1, null, Context.NONE);
    }
}
