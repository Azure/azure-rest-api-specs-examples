import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections List. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_listprivateendpointconnections.json
     */
    /**
     * Sample code: PrivateEndpointConnections_List.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void privateEndpointConnectionsList(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.privateEndpointConnections().listWithResponse("myResourceGroup", "testHub", Context.NONE);
    }
}
