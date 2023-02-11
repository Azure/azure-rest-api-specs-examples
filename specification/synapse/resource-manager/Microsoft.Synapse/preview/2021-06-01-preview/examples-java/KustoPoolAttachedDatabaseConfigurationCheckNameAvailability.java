import com.azure.resourcemanager.synapse.models.DatabaseCheckNameRequest;
import com.azure.resourcemanager.synapse.models.Type;

/** Samples for KustoPoolChildResource CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationCheckNameAvailability.json
     */
    /**
     * Sample code: KustoPoolAttachedDatabaseConfigurationCheckNameAvailability.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolAttachedDatabaseConfigurationCheckNameAvailability(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolChildResources()
            .checkNameAvailabilityWithResponse(
                "kustorptest",
                "kustoclusterrptest4",
                "kustorptest",
                new DatabaseCheckNameRequest()
                    .withName("adc1")
                    .withType(Type.MICROSOFT_SYNAPSE_WORKSPACES_KUSTO_POOLS_ATTACHED_DATABASE_CONFIGURATIONS),
                com.azure.core.util.Context.NONE);
    }
}
