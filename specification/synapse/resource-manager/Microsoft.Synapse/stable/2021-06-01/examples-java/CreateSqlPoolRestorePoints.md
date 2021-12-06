Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.CreateSqlPoolRestorePointDefinition;

/** Samples for SqlPoolRestorePoints Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateSqlPoolRestorePoints.json
     */
    /**
     * Sample code: Creates Sql pool restore point.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createsSqlPoolRestorePoint(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolRestorePoints()
            .create(
                "Default-SQL-SouthEastAsia",
                "testserver",
                "testDatabase",
                new CreateSqlPoolRestorePointDefinition().withRestorePointLabel("mylabel"),
                Context.NONE);
    }
}
```
