import com.azure.core.util.Context;

/** Samples for PrivateEndpointConnection Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeletePrivateEndpointConnection.json
     */
    /**
     * Sample code: ApiManagementDeletePrivateEndpointConnection.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeletePrivateEndpointConnection(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .privateEndpointConnections()
            .delete("rg1", "apimService1", "privateEndpointConnectionName", Context.NONE);
    }
}
