Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlPoolOperations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolOperations.json
     */
    /**
     * Sample code: List the Sql Analytics pool management operations.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listTheSqlAnalyticsPoolManagementOperations(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolOperations().list("sqlcrudtest-7398", "sqlcrudtest-4645", "testdb", Context.NONE);
    }
}
```
