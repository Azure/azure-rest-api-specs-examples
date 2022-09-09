import com.azure.resourcemanager.agrifood.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.agrifood.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/agrifood/resource-manager/Microsoft.AgFoodPlatform/preview/2021-09-01-preview/examples/PrivateEndpointConnections_CreateOrUpdate.json
     */
    /**
     * Sample code: PrivateEndpointConnections_CreateOrUpdate.
     *
     * @param manager Entry point to AgriFoodManager.
     */
    public static void privateEndpointConnectionsCreateOrUpdate(
        com.azure.resourcemanager.agrifood.AgriFoodManager manager) {
        manager
            .privateEndpointConnections()
            .define("privateEndpointConnectionName")
            .withExistingFarmBeat("examples-rg", "examples-farmbeatsResourceName")
            .withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState()
                    .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Approved by johndoe@contoso.com"))
            .create();
    }
}
