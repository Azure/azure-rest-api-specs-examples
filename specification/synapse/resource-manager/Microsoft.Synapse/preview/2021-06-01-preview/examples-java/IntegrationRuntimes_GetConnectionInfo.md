Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for IntegrationRuntimeConnectionInfos Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_GetConnectionInfo.json
     */
    /**
     * Sample code: Get connection info.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getConnectionInfo(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimeConnectionInfos()
            .getWithResponse("exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", Context.NONE);
    }
}
```
