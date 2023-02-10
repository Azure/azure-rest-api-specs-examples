/** Samples for KustoPools ListSkusByResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsListResourceSkus.json
     */
    /**
     * Sample code: KustoPoolsListResourceSkus.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolsListResourceSkus(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPools()
            .listSkusByResource(
                "synapseWorkspaceName", "kustoclusterrptest4", "kustorptest", com.azure.core.util.Context.NONE);
    }
}
