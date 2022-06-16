import com.azure.core.util.Context;

/** Samples for GatewayHostnameConfiguration Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetGatewayHostnameConfiguration.json
     */
    /**
     * Sample code: ApiManagementGetGatewayHostnameConfiguration.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetGatewayHostnameConfiguration(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.gatewayHostnameConfigurations().getWithResponse("rg1", "apimService1", "gw1", "default", Context.NONE);
    }
}
