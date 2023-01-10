import com.azure.core.util.Context;

/** Samples for OnlineDeployments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineDeployment/delete.json
     */
    /**
     * Sample code: Delete Online Deployment.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteOnlineDeployment(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineDeployments().delete("testrg123", "workspace123", "testEndpoint", "testDeployment", Context.NONE);
    }
}
