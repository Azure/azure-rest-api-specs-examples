Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.ResourceMoveDefinition;

/** Samples for SqlPools Rename. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/RenameSqlPool.json
     */
    /**
     * Sample code: Rename a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void renameASQLAnalyticsPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPools()
            .renameWithResponse(
                "Default-SQL-SouthEastAsia",
                "testsvr",
                "testdb",
                new ResourceMoveDefinition()
                    .withId(
                        "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Synapse/workspaces/testsvr/sqlPools/newtestdb"),
                Context.NONE);
    }
}
```
