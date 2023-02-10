import com.azure.resourcemanager.synapse.fluent.models.DataConnectionValidationInner;
import com.azure.resourcemanager.synapse.models.EventHubDataConnection;

/** Samples for KustoPoolDataConnections DataConnectionValidation. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionValidation.json
     */
    /**
     * Sample code: KustoPoolDataConnectionValidation.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDataConnectionValidation(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDataConnections()
            .dataConnectionValidation(
                "kustorptest",
                "kustorptest",
                "kustoclusterrptest4",
                "KustoDatabase8",
                new DataConnectionValidationInner()
                    .withDataConnectionName("DataConnections8")
                    .withProperties(new EventHubDataConnection()),
                com.azure.core.util.Context.NONE);
    }
}
