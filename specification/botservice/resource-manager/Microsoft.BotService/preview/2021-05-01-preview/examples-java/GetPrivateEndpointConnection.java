import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/botservice/resource-manager/Microsoft.BotService/preview/2021-05-01-preview/examples/GetPrivateEndpointConnection.json
     */
    /**
     * Sample code: Get Private Endpoint Connection.
     *
     * @param manager Entry point to BotServiceManager.
     */
    public static void getPrivateEndpointConnection(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager
            .privateEndpointConnections()
            .getWithResponse("res6977", "sto2527", "{privateEndpointConnectionName}", Context.NONE);
    }
}
