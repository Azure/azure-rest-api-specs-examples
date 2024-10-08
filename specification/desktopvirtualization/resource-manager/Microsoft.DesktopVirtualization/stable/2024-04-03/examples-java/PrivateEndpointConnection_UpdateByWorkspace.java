
import com.azure.resourcemanager.desktopvirtualization.models.PrivateEndpointConnection;
import com.azure.resourcemanager.desktopvirtualization.models.PrivateEndpointServiceConnectionStatus;
import com.azure.resourcemanager.desktopvirtualization.models.PrivateLinkServiceConnectionState;

/**
 * Samples for PrivateEndpointConnections UpdateByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * PrivateEndpointConnection_UpdateByWorkspace.json
     */
    /**
     * Sample code: PrivateEndpointConnection_UpdateByWorkspace.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void privateEndpointConnectionUpdateByWorkspace(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.privateEndpointConnections().updateByWorkspaceWithResponse("resourceGroup1", "workspace1",
            "workspace1.377103f1-5179-4bdf-8556-4cdd3207cc5b",
            new PrivateEndpointConnection().withPrivateLinkServiceConnectionState(
                new PrivateLinkServiceConnectionState().withStatus(PrivateEndpointServiceConnectionStatus.APPROVED)
                    .withDescription("Approved by admin@consoto.com").withActionsRequired("None")),
            com.azure.core.util.Context.NONE);
    }
}
