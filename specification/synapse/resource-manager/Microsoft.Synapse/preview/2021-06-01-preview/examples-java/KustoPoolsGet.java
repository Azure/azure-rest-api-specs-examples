import com.azure.core.util.Context;

/** Samples for KustoPools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsGet.json
     */
    /**
     * Sample code: kustoPoolsGet.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolsGet(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPools()
            .getWithResponse("synapseWorkspaceName", "kustoclusterrptest4", "kustorptest", Context.NONE);
    }
}
