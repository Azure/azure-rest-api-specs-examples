/** Samples for KustoPoolDataConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsDelete.json
     */
    /**
     * Sample code: KustoPoolDataConnectionsDelete.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDataConnectionsDelete(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDataConnections()
            .delete(
                "kustorptest",
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "KustoDatabase8",
                "kustoeventhubconnection1",
                com.azure.core.util.Context.NONE);
    }
}
