Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.5/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.ReadWriteDatabase;
import java.time.Duration;

/** Samples for KustoPoolDatabases Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasesUpdate.json
     */
    /**
     * Sample code: KustoPoolDatabasesUpdate.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDatabasesUpdate(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDatabases()
            .update(
                "kustorptest",
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "KustoDatabase8",
                new ReadWriteDatabase().withSoftDeletePeriod(Duration.parse("P1D")),
                Context.NONE);
    }
}
```
