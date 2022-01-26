Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for RestorableDroppedSqlPools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/RestorableDroppedSqlPoolGet.json
     */
    /**
     * Sample code: Get a restorable dropped Sql pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getARestorableDroppedSqlPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .restorableDroppedSqlPools()
            .getWithResponse(
                "restorabledroppeddatabasetest-1257",
                "restorabledroppeddatabasetest-2389",
                "restorabledroppeddatabasetest-7654,131403269876900000",
                Context.NONE);
    }
}
```
