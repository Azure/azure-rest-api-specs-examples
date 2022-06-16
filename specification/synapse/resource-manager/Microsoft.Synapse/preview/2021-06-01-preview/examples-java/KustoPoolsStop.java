import com.azure.core.util.Context;

/** Samples for KustoPools Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsStop.json
     */
    /**
     * Sample code: kustoPoolsStop.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolsStop(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.kustoPools().stop("kustorptest", "kustoclusterrptest4", "kustorptest", Context.NONE);
    }
}
