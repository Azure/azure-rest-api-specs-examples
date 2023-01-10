import com.azure.core.util.Context;
import com.azure.resourcemanager.machinelearning.models.EndpointComputeType;
import com.azure.resourcemanager.machinelearning.models.OrderString;

/** Samples for OnlineEndpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/OnlineEndpoint/list.json
     */
    /**
     * Sample code: List Online Endpoint.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listOnlineEndpoint(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
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
