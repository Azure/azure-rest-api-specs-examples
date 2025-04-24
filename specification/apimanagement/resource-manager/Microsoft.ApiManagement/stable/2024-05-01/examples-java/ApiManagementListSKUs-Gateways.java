
/**
 * Samples for ApiManagementGatewaySkus ListAvailableSkus.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListSKUs-Gateways.json
     */
    /**
     * Sample code: ApiManagementListSKUs-Gateways.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListSKUsGateways(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementGatewaySkus().listAvailableSkus("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
