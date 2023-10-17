/** Samples for PrivateEndpointConnections GetByHostPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/PrivateEndpointConnection_GetByHostPool.json
     */
    /**
     * Sample code: PrivateEndpointConnection_GetByHostPool.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void privateEndpointConnectionGetByHostPool(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .privateEndpointConnections()
            .getByHostPoolWithResponse(
                "resourceGroup1",
                "hostPool1",
                "hostPool1.377103f1-5179-4bdf-8556-4cdd3207cc5b",
                com.azure.core.util.Context.NONE);
    }
}
