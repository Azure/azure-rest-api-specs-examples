
/**
 * Samples for ApiGateway ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListGatewaysBySubscriptionAndResourceGroup.json
     */
    /**
     * Sample code: ApiManagementListGatewaysBySubscriptionAndResourceGroup.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListGatewaysBySubscriptionAndResourceGroup(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiGateways().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
