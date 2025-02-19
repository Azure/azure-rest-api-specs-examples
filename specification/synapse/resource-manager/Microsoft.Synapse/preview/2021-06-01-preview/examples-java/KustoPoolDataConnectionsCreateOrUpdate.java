
import com.azure.resourcemanager.synapse.models.EventHubDataConnection;

/**
 * Samples for KustoPoolDataConnections CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/
     * KustoPoolDataConnectionsCreateOrUpdate.json
     */
    /**
     * Sample code: KustoPoolDataConnectionsCreateOrUpdate.json.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void
        kustoPoolDataConnectionsCreateOrUpdateJson(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.kustoPoolDataConnections().createOrUpdate("kustorptest", "synapseWorkspaceName", "kustoclusterrptest4",
            "KustoDatabase8", "DataConnections8",
            new EventHubDataConnection().withLocation("westus").withEventHubResourceId(
                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.EventHub/namespaces/eventhubTestns1/eventhubs/eventhubTest1")
                .withConsumerGroup("testConsumerGroup1"),
            com.azure.core.util.Context.NONE);
    }
}
