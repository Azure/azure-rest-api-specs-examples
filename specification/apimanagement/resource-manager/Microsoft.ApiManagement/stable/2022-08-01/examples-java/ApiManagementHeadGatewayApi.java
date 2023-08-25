/** Samples for GatewayApi GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadGatewayApi.json
     */
    /**
     * Sample code: ApiManagementHeadGatewayApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadGatewayApi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayApis()
            .getEntityTagWithResponse("rg1", "apimService1", "gw1", "api1", com.azure.core.util.Context.NONE);
    }
}
