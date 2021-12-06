Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlPoolReplicationLinks GetByName. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolReplicationLinks_GetByName.json
     */
    /**
     * Sample code: Lists a Sql Analytic pool's replication links.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listsASqlAnalyticPoolSReplicationLinks(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolReplicationLinks()
            .getByNameWithResponse(
                "sqlcrudtest-4799", "sqlcrudtest-6440", "testdb", "5b301b68-03f6-4b26-b0f4-73ebb8634238", Context.NONE);
    }
}
```
