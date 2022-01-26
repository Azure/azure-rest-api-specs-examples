Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlPoolSchemas List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolSchema.json
     */
    /**
     * Sample code: List the schema in a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listTheSchemaInASQLAnalyticsPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolSchemas().list("myRG", "serverName", "myDatabase", null, Context.NONE);
    }
}
```
