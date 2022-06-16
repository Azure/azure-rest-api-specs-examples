import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections ListByHostPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/PrivateEndpointConnection_ListByHostPool.json
     */
    /**
     * Sample code: PrivateEndpointConnection_ListByHostPool.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void privateEndpointConnectionListByHostPool(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.privateEndpointConnections().listByHostPool("resourceGroup1", "hostPool1", Context.NONE);
    }
}
