Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-machinelearning_1.0.0-beta.1/sdk/machinelearning/azure-resourcemanager-machinelearning/README.md) on how to add the SDK to your project and authenticate.

```java
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
```
