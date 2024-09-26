
/**
 * Samples for PrivateEndpointConnections DeleteByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * PrivateEndpointConnection_DeleteByWorkspace.json
     */
    /**
     * Sample code: PrivateEndpointConnection_DeleteByWorkspace.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void privateEndpointConnectionDeleteByWorkspace(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.privateEndpointConnections().deleteByWorkspaceWithResponse("resourceGroup1", "workspace1",
            "workspace1.377103f1-5179-4bdf-8556-4cdd3207cc5b", com.azure.core.util.Context.NONE);
    }
}
