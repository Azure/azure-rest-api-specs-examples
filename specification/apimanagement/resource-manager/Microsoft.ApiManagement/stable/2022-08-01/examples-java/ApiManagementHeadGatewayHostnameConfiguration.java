/** Samples for GatewayHostnameConfiguration GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadGatewayHostnameConfiguration.json
     */
    /**
     * Sample code: ApiManagementHeadGatewayHostnameConfiguration.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadGatewayHostnameConfiguration(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayHostnameConfigurations()
            .getEntityTagWithResponse("rg1", "apimService1", "gw1", "default", com.azure.core.util.Context.NONE);
    }
}
