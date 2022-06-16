import com.azure.core.util.Context;

/** Samples for OnlineDeployments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineDeployment/list.json
     */
    /**
     * Sample code: List Online Deployments.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listOnlineDeployments(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .onlineDeployments()
            .list("test-rg", "my-aml-workspace", "testEndpointName", "string", 1, null, Context.NONE);
    }
}
