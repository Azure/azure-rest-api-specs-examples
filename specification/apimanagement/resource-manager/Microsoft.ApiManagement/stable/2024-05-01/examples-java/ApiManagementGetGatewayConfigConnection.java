
/**
 * Samples for ApiGatewayConfigConnection Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetGatewayConfigConnection.json
     */
    /**
     * Sample code: ApiManagementGetGatewayConfigConnection.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetGatewayConfigConnection(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiGatewayConfigConnections().getWithResponse("rg1", "standard-gw-01", "gcc-01",
            com.azure.core.util.Context.NONE);
    }
}
