/** Samples for GatewayHostnameConfiguration Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteGatewayHostnameConfiguration.json
     */
    /**
     * Sample code: ApiManagementDeleteGatewayHostnameConfiguration.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteGatewayHostnameConfiguration(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayHostnameConfigurations()
            .deleteWithResponse("rg1", "apimService1", "gw1", "default", "*", com.azure.core.util.Context.NONE);
    }
}
