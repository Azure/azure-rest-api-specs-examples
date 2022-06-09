```java
import com.azure.core.util.Context;

/** Samples for IntegrationRuntimeCredentials Sync. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/IntegrationRuntimes_SyncCredentials.json
     */
    /**
     * Sample code: Sync credentials.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void syncCredentials(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .integrationRuntimeCredentials()
            .syncWithResponse("exampleResourceGroup", "exampleWorkspace", "exampleIntegrationRuntime", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
