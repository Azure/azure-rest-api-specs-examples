import com.azure.resourcemanager.iothub.fluent.models.PrivateEndpointConnectionInner;
import com.azure.resourcemanager.iothub.models.PrivateEndpointConnectionProperties;
import com.azure.resourcemanager.iothub.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.iothub.models.PrivateLinkServiceConnectionStatus;

/** Samples for PrivateEndpointConnections Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2023-06-30-preview/examples/iothub_updateprivateendpointconnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_Update.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void privateEndpointConnectionUpdate(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .privateEndpointConnections()
            .update(
                "myResourceGroup",
                "testHub",
                "myPrivateEndpointConnection",
                new PrivateEndpointConnectionInner()
                    .withProperties(
                        new PrivateEndpointConnectionProperties()
                            .withPrivateLinkServiceConnectionState(
                                new PrivateLinkServiceConnectionState()
                                    .withStatus(PrivateLinkServiceConnectionStatus.APPROVED)
                                    .withDescription("Approved by johndoe@contoso.com"))),
                com.azure.core.util.Context.NONE);
    }
}
