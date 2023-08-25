/** Samples for GatewayApi Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteGatewayApi.json
     */
    /**
     * Sample code: ApiManagementDeleteGatewayApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteGatewayApi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayApis()
            .deleteWithResponse("rg1", "apimService1", "gw1", "echo-api", com.azure.core.util.Context.NONE);
    }
}
