/** Samples for KustoPoolDatabases ListByKustoPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoDatabasesListByKustoPool.json
     */
    /**
     * Sample code: KustoDatabasesListByKustoPool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoDatabasesListByKustoPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDatabases()
            .listByKustoPool(
                "kustorptest", "synapseWorkspaceName", "kustoclusterrptest4", com.azure.core.util.Context.NONE);
    }
}
