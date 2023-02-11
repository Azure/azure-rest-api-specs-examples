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
                com.azure.core.util.Context.NONE);
    }
}
