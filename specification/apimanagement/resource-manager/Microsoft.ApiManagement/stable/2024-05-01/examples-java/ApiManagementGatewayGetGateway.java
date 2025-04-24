
/**
 * Samples for ApiGateway GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGatewayGetGateway.json
     */
    /**
     * Sample code: ApiManagementGatewayGetGateway.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGatewayGetGateway(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiGateways().getByResourceGroupWithResponse("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
