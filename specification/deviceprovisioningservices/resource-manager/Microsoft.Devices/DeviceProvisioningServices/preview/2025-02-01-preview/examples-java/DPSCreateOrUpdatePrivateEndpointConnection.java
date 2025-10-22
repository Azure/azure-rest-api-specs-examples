
import com.azure.resourcemanager.deviceprovisioningservices.models.PrivateEndpointConnectionProperties;
import com.azure.resourcemanager.deviceprovisioningservices.models.PrivateLinkServiceConnectionState;
import com.azure.resourcemanager.deviceprovisioningservices.models.PrivateLinkServiceConnectionStatus;

/**
 * Samples for IotDpsResource CreateOrUpdatePrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSCreateOrUpdatePrivateEndpointConnection.json
     */
    /**
     * Sample code: PrivateEndpointConnection_CreateOrUpdate.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void privateEndpointConnectionCreateOrUpdate(
        com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().definePrivateEndpointConnection("myPrivateEndpointConnection")
            .withExistingProvisioningService("myResourceGroup", "myFirstProvisioningService")
            .withProperties(new PrivateEndpointConnectionProperties().withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateLinkServiceConnectionStatus.APPROVED)
                    .withDescription("Approved by johndoe@contoso.com")))
            .create();
    }
}
