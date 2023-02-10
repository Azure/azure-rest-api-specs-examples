/** Samples for KustoPoolDataConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsGet.json
     */
    /**
     * Sample code: KustoPoolDataConnectionsGet.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDataConnectionsGet(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDataConnections()
            .getWithResponse(
                "kustorptest",
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "KustoDatabase8",
                "DataConnections8",
                com.azure.core.util.Context.NONE);
    }
}
