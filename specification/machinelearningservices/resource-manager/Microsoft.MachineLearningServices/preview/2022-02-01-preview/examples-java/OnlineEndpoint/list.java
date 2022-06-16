import com.azure.core.util.Context;
import com.azure.resourcemanager.machinelearning.models.EndpointComputeType;
import com.azure.resourcemanager.machinelearning.models.OrderString;

/** Samples for OnlineEndpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/OnlineEndpoint/list.json
     */
    /**
     * Sample code: List Online Endpoint.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listOnlineEndpoint(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .onlineEndpoints()
            .list(
                "test-rg",
                "my-aml-workspace",
                "string",
                1,
                EndpointComputeType.MANAGED,
                null,
                "string",
                "string",
                OrderString.CREATED_AT_DESC,
                Context.NONE);
    }
}
