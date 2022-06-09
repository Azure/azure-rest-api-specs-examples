```java
import com.azure.core.util.Context;

/** Samples for SqlPoolRestorePoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolRestorePointsDelete.json
     */
    /**
     * Sample code: Deletes a restore point.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deletesARestorePoint(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolRestorePoints()
            .deleteWithResponse("Default-SQL-SouthEastAsia", "testws", "testpool", "131546477590000000", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
