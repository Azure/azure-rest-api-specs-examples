import com.azure.core.util.Context;

/** Samples for KustoPoolDatabases Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasesGet.json
     */
    /**
     * Sample code: KustoPoolDatabasesGet.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDatabasesGet(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDatabases()
            .getWithResponse(
                "kustorptest", "synapseWorkspaceName", "kustoclusterrptest4", "KustoDatabase8", Context.NONE);
    }
}
