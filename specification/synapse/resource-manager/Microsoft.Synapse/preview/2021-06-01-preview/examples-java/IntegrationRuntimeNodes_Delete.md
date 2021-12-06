Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for IntegrationRuntimeNodes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimeNodes_Delete.json
     */
    /**
     * Sample code: Delete integration runtime node.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deleteIntegrationRuntimeNode(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimeNodes()
            .deleteWithResponse(
                "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", "Node_1", Context.NONE);
    }
}
```
