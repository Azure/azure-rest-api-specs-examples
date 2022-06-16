import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_deleteprivateendpointconnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_Delete.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void privateEndpointConnectionDelete(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .privateEndpointConnections()
            .delete("myResourceGroup", "testHub", "myPrivateEndpointConnection", Context.NONE);
    }
}
