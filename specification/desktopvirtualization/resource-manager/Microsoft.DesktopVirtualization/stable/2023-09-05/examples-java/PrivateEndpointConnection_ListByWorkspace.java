/** Samples for PrivateEndpointConnections ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/PrivateEndpointConnection_ListByWorkspace.json
     */
    /**
     * Sample code: PrivateEndpointConnection_ListByWorkspace.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void privateEndpointConnectionListByWorkspace(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .privateEndpointConnections()
            .listByWorkspace("resourceGroup1", "workspace1", com.azure.core.util.Context.NONE);
    }
}
