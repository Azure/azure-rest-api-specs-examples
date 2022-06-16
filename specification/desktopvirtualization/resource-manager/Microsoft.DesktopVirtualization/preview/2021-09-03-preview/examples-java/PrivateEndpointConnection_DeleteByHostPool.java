import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnections DeleteByHostPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/PrivateEndpointConnection_DeleteByHostPool.json
     */
    /**
     * Sample code: PrivateEndpointConnection_DeleteByHostPool.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void privateEndpointConnectionDeleteByHostPool(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .privateEndpointConnections()
            .deleteByHostPoolWithResponse(
                "resourceGroup1", "hostPool1", "hostPool1.377103f1-5179-4bdf-8556-4cdd3207cc5b", Context.NONE);
    }
}
