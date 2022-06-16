import com.azure.core.util.Context;
import com.azure.resourcemanager.desktopvirtualization.models.PrivateEndpointConnection;
import com.azure.resourcemanager.desktopvirtualization.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.desktopvirtualization.models.PrivateLinkServiceConnectionState;

/** Samples for PrivateEndpointConnections UpdateByHostPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/PrivateEndpointConnection_UpdateByHostPool.json
     */
    /**
     * Sample code: PrivateEndpointConnection_UpdateByHostPool.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void privateEndpointConnectionUpdateByHostPool(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .privateEndpointConnections()
            .updateByHostPoolWithResponse(
                "resourceGroup1",
                "hostPool1",
                "hostPool1.377103f1-5179-4bdf-8556-4cdd3207cc5b",
                new PrivateEndpointConnection()
                    .withPrivateLinkServiceConnectionState(
                        new PrivateLinkServiceConnectionState()
                            .withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                            .withDescription("Approved by admin@consoto.com")
                            .withActionsRequired("None")),
                Context.NONE);
    }
}
