import com.azure.resourcemanager.synapse.models.KustoPoolCheckNameRequest;

/** Samples for KustoPools CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolsCheckNameAvailability.json
     */
    /**
     * Sample code: KustoPoolsCheckNameAvailability.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolsCheckNameAvailability(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPools()
            .checkNameAvailabilityWithResponse(
                "westus",
                new KustoPoolCheckNameRequest().withName("kustoclusterrptest4"),
                com.azure.core.util.Context.NONE);
    }
}
