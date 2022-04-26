Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.IntegrationRuntimeAutoUpdate;
import com.azure.resourcemanager.synapse.models.IntegrationRuntimeResource;

/** Samples for IntegrationRuntimes Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_Update.json
     */
    /**
     * Sample code: Update integration runtime.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void updateIntegrationRuntime(com.azure.resourcemanager.synapse.SynapseManager manager) {
        IntegrationRuntimeResource resource =
            manager
                .integrationRuntimes()
                .getWithResponse(
                    "exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", null, Context.NONE)
                .getValue();
        resource.update().withAutoUpdate(IntegrationRuntimeAutoUpdate.OFF).withUpdateDelayOffset("\"PT3H\"").apply();
    }
}
```
