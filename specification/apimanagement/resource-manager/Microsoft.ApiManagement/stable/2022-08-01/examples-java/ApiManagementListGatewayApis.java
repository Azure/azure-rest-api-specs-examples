/** Samples for GatewayApi ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGatewayApis.json
     */
    /**
     * Sample code: ApiManagementListGatewayApis.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListGatewayApis(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .gatewayApis()
            .listByService("rg1", "apimService1", "gw1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
