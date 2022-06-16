import com.azure.core.util.Context;

/** Samples for OnlineDeployments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineDeployment/delete.json
     */
    /**
     * Sample code: Delete Online Deployment.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteOnlineDeployment(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.onlineDeployments().delete("testrg123", "workspace123", "testEndpoint", "testDeployment", Context.NONE);
    }
}
