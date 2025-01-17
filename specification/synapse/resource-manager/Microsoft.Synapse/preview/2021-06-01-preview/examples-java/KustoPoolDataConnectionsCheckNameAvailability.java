
import com.azure.resourcemanager.synapse.models.DataConnectionCheckNameRequest;

/**
 * Samples for KustoPoolDataConnections CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/
     * KustoPoolDataConnectionsCheckNameAvailability.json
     */
    /**
     * Sample code: KustoPoolDataConnectionsCheckNameAvailability.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void
        kustoPoolDataConnectionsCheckNameAvailability(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.kustoPoolDataConnections().checkNameAvailabilityWithResponse("kustorptest", "synapseWorkspaceName",
            "kustoclusterrptest4", "Kustodatabase8", new DataConnectionCheckNameRequest().withName("DataConnections8"),
            com.azure.core.util.Context.NONE);
    }
}
