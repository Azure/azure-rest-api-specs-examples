import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections GetByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/PrivateEndpointConnection_GetByWorkspace.json
     */
    /**
     * Sample code: PrivateEndpointConnection_GetByWorkspace.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void privateEndpointConnectionGetByWorkspace(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .privateEndpointConnections()
            .getByWorkspaceWithResponse(
                "resourceGroup1", "workspace1", "workspace1.377103f1-5179-4bdf-8556-4cdd3207cc5b", Context.NONE);
    }
}
