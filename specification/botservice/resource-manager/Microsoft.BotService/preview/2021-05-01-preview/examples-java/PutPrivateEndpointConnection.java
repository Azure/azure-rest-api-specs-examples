import com.azure.resourcemanager.botservice.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.botservice.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/PutPrivateEndpointConnection.json
     */
    /**
     * Sample code: Put Private Endpoint Connection.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void putPrivateEndpointConnection(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .privateEndpointConnections()
            .define("{privateEndpointConnectionName}")
            .withExistingBotService("res7687", "sto9699")
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState()
                    .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Auto-Approved"))
            .create();
    }
}
