Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for IntegrationRuntimeNodeIpAddressOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimeNodes_GetIpAddress.json
     */
    /**
     * Sample code: Get integration runtime node IP address.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getIntegrationRuntimeNodeIPAddress(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimeNodeIpAddressOperations()
            .getWithResponse(
                "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", "Node_1", Context.NONE);
    }
}
```
