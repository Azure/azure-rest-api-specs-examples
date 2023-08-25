import com.azure.resourcemanager.apimanagement.models.PrivateEndpointConnectionRequestProperties;
import com.azure.resourcemanager.apimanagement.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.apimanagement.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnection CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementApproveOrRejectPrivateEndpointConnection.json
     */
    /**
     * Sample code: ApiManagementApproveOrRejectPrivateEndpointConnection.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementApproveOrRejectPrivateEndpointConnection(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .privateEndpointConnections()
            .define("privateEndpointConnectionName")
            .withExistingService("rg1", "apimService1")
            .withProperties(
                new PrivateEndpointConnectionRequestProperties()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                            .withDescription("The Private Endpoint Connection is approved.")))
            .create();
    }
}
