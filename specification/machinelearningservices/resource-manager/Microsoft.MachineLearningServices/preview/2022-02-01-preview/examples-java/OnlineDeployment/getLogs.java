import com.azure.core.util.Context;
import com.azure.resourcemanager.machinelearning.models.ContainerType;
import com.azure.resourcemanager.machinelearning.models.DeploymentLogsRequest;

/** Samples for OnlineDeployments GetLogs. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineDeployment/getLogs.json
     */
    /**
     * Sample code: Get Online Deployment Logs.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getOnlineDeploymentLogs(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .onlineDeployments()
            .getLogsWithResponse(
                "testrg123",
                "workspace123",
                "testEndpoint",
                "testDeployment",
                new DeploymentLogsRequest().withContainerType(ContainerType.STORAGE_INITIALIZER).withTail(0),
                Context.NONE);
    }
}
