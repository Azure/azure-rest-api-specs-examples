/** Samples for GatewayHostnameConfiguration ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGatewayHostnameConfigurations.json
     */
    /**
     * Sample code: ApiManagementListGatewayHostnameConfigurations.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListGatewayHostnameConfigurations(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayHostnameConfigurations()
            .listByService("rg1", "apimService1", "gw1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
