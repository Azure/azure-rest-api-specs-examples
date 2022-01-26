Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.fluent.models.MetadataSyncConfigInner;

/** Samples for SqlPoolMetadataSyncConfigs Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateSqlPoolMetadataSyncConfig.json
     */
    /**
     * Sample code: Set metadata sync config for a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void setMetadataSyncConfigForASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolMetadataSyncConfigs()
            .createWithResponse(
                "ExampleResourceGroup",
                "ExampleWorkspace",
                "ExampleSqlPool",
                new MetadataSyncConfigInner().withEnabled(true),
                Context.NONE);
    }
}
```
