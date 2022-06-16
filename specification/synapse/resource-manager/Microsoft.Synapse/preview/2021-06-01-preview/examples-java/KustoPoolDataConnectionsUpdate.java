import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.models.EventHubDataConnection;

/** Samples for KustoPoolDataConnections Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolDataConnectionsUpdate.json
     */
    /**
     * Sample code: KustoPoolDataConnectionsUpdate.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDataConnectionsUpdate(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolDataConnections()
            .update(
                "kustorptest",
                "synapseWorkspaceName",
                "kustoclusterrptest4",
                "KustoDatabase8",
                "DataConnections8",
                new EventHubDataConnection()
                    .withLocation("westus")
                    .withEventHubResourceId(
                        "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1")
                    .withConsumerGroup("testConsumerGroup1"),
                Context.NONE);
    }
}
