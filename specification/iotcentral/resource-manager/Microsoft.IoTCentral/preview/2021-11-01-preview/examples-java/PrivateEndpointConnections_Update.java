import com.azure.resourcemanager.iotcentral.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.iotcentral.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateEndpointConnections_Update.json
     */
    /**
     * Sample code: PrivateEndpointConnections_Update.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void privateEndpointConnectionsUpdate(
        com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager
            .privateEndpointConnections()
            .define("myIoTCentralAppEndpoint")
            .withExistingIotApp("resRg", "myIoTCentralApp")
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState()
                    .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Auto-approved")
                    .withActionsRequired("None"))
            .create();
    }
}
