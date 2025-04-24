
/**
 * Samples for Gateway InvalidateDebugCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGatewayInvalidateDebugCredentials.json
     */
    /**
     * Sample code: ApiManagementGatewayInvalidateDebugCredentials.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGatewayInvalidateDebugCredentials(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.gateways().invalidateDebugCredentialsWithResponse("rg1", "apimService1", "gw1",
            com.azure.core.util.Context.NONE);
    }
}
