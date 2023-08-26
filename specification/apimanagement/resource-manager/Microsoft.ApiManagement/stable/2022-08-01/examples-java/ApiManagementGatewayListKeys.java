/** Samples for Gateway ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGatewayListKeys.json
     */
    /**
     * Sample code: ApiManagementGatewayListKeys.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGatewayListKeys(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.gateways().listKeysWithResponse("rg1", "apimService1", "gw1", com.azure.core.util.Context.NONE);
    }
}
