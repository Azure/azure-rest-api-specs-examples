/** Samples for PrivateEndpointConnection GetByName. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetPrivateEndpointConnection.json
     */
    /**
     * Sample code: ApiManagementGetPrivateEndpointConnection.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetPrivateEndpointConnection(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .privateEndpointConnections()
            .getByNameWithResponse(
                "rg1", "apimService1", "privateEndpointConnectionName", com.azure.core.util.Context.NONE);
    }
}
