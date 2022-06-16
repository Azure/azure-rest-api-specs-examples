import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/DeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: Delete Private Endpoint Connection.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void deletePrivateEndpointConnection(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .privateEndpointConnections()
            .deleteWithResponse("res6977", "sto2527", "{privateEndpointConnectionName}", Context.NONE);
    }
}
