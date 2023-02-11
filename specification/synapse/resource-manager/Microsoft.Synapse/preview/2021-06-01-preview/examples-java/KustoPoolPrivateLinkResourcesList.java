/** Samples for KustoPoolPrivateLinkResourcesOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrivateLinkResourcesList.json
     */
    /**
     * Sample code: KustoPoolPrivateLinkResourcesList.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolPrivateLinkResourcesList(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolPrivateLinkResourcesOperations()
            .list("DP-900", "synapse-ws-ebi-data", "dataexplorerpool900", com.azure.core.util.Context.NONE);
    }
}
