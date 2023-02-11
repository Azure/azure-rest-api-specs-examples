/** Samples for KustoPools Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsDelete.json
     */
    /**
     * Sample code: kustoPoolsDelete.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolsDelete(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPools()
            .delete("kustorptest", "kustorptest", "kustoclusterrptest4", com.azure.core.util.Context.NONE);
    }
}
