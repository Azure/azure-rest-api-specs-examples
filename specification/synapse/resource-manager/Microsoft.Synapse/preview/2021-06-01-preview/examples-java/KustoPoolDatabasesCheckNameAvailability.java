import com.azure.resourcemanager.synapse.models.DatabaseCheckNameRequest;
import com.azure.resourcemanager.synapse.models.Type;

/** Samples for KustoPoolChildResource CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDatabasesCheckNameAvailability.json
     */
    /**
     * Sample code: KustoPoolDatabasesCheckNameAvailability.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDatabasesCheckNameAvailability(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolChildResources()
            .checkNameAvailabilityWithResponse(
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "kustorptest",
                new DatabaseCheckNameRequest()
                    .withName("database1")
                    .withType(Type.MICROSOFT_SYNAPSE_WORKSPACES_KUSTO_POOLS_DATABASES),
                com.azure.core.util.Context.NONE);
    }
}
