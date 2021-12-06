Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for RestorableDroppedSqlPools ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/RestorableDroppedSqlpoolList.json
     */
    /**
     * Sample code: Get list of restorable dropped Sql pools.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getListOfRestorableDroppedSqlPools(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .restorableDroppedSqlPools()
            .listByWorkspace("restorabledroppeddatabasetest-1349", "restorabledroppeddatabasetest-1840", Context.NONE);
    }
}
```
