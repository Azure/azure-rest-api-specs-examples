/** Samples for KustoPoolDataConnections ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsListByDatabase.json
     */
    /**
     * Sample code: KustoPoolDataConnectionsListByDatabase.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDataConnectionsListByDatabase(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDataConnections()
            .listByDatabase(
                "kustorptest",
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "KustoDatabase8",
                com.azure.core.util.Context.NONE);
    }
}
