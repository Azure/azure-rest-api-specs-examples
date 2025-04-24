
/**
 * Samples for ApiGateway Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGatewayDeleteGateway.json
     */
    /**
     * Sample code: ApiManagementGatewayDeleteGateway.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGatewayDeleteGateway(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiGateways().delete("rg1", "example-gateway", com.azure.core.util.Context.NONE);
    }
}
