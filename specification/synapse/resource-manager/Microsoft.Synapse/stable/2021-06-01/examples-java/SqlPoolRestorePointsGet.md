Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SqlPoolRestorePoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolRestorePointsGet.json
     */
    /**
     * Sample code: Gets a Sql pool restore point.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getsASqlPoolRestorePoint(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolRestorePoints()
            .getWithResponse("Default-SQL-SouthEastAsia", "testws", "testpool", "131546477590000000", Context.NONE);
    }
}
```
