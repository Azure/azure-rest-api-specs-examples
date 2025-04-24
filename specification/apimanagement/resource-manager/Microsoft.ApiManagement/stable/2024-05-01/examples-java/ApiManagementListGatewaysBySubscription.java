
/**
 * Samples for ApiGateway List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListGatewaysBySubscription.json
     */
    /**
     * Sample code: ApiManagementListGatewaysBySubscription.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListGatewaysBySubscription(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiGateways().list(com.azure.core.util.Context.NONE);
    }
}
