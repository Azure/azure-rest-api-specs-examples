
/**
 * Samples for PrivateEndpointConnections ListByHostPool.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * PrivateEndpointConnection_ListByHostPool.json
     */
    /**
     * Sample code: PrivateEndpointConnection_ListByHostPool.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void privateEndpointConnectionListByHostPool(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.privateEndpointConnections().listByHostPool("resourceGroup1", "hostPool1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
