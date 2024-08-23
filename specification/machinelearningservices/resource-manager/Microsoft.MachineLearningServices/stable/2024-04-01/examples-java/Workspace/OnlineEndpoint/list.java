
import com.azure.resourcemanager.machinelearning.models.EndpointComputeType;
import com.azure.resourcemanager.machinelearning.models.OrderString;

/**
 * Samples for OnlineEndpoints List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/OnlineEndpoint/list.json
     */
    /**
     * Sample code: List Workspace Online Endpoint.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceOnlineEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.onlineEndpoints().list("test-rg", "my-aml-workspace", "string", 1, EndpointComputeType.MANAGED, null,
            "string", "string", OrderString.CREATED_AT_DESC, com.azure.core.util.Context.NONE);
    }
}
