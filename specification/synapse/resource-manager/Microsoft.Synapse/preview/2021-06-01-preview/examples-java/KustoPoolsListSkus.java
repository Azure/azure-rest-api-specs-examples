import com.azure.core.util.Context;

/** Samples for KustoPools List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsListSkus.json
     */
    /**
     * Sample code: KustoPoolsListSkus.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolsListSkus(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.kustoPools().list(Context.NONE);
    }
}
