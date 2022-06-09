```java
import com.azure.core.util.Context;

/** Samples for SqlPoolReplicationLinks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolReplicationLinks.json
     */
    /**
     * Sample code: Lists a Sql Analytic pool's replication links.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listsASqlAnalyticPoolSReplicationLinks(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolReplicationLinks().list("sqlcrudtest-4799", "sqlcrudtest-6440", "testdb", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
