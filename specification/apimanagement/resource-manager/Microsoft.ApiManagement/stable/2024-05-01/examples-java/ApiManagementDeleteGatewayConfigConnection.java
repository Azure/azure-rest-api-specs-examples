
/**
 * Samples for ApiGatewayConfigConnection Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteGatewayConfigConnection.json
     */
    /**
     * Sample code: ApiManagementGatewayDeleteGateway.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGatewayDeleteGateway(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiGatewayConfigConnections().delete("rg1", "standard-gw-01", "gcc-01", "*",
            com.azure.core.util.Context.NONE);
    }
}
