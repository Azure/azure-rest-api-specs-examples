Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlPoolColumns Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolColumnGet.json
     */
    /**
     * Sample code: Get database column.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getDatabaseColumn(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolColumns()
            .getWithResponse("myRG", "serverName", "myDatabase", "dbo", "table1", "column1", Context.NONE);
    }
}
```
