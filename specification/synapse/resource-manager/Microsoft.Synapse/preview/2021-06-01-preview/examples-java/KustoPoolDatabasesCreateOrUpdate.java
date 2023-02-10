import com.azure.resourcemanager.synapse.models.ReadWriteDatabase;
import java.time.Duration;

/** Samples for KustoPoolDatabases CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasesCreateOrUpdate.json
     */
    /**
     * Sample code: KustoPoolDatabasesCreateOrUpdate.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDatabasesCreateOrUpdate(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDatabases()
            .createOrUpdate(
                "kustorptest",
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "KustoDatabase8",
                new ReadWriteDatabase().withLocation("westus").withSoftDeletePeriod(Duration.parse("P1D")),
                com.azure.core.util.Context.NONE);
    }
}
