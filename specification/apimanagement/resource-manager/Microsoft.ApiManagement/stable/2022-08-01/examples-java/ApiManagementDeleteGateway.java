/** Samples for Gateway Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteGateway.json
     */
    /**
     * Sample code: ApiManagementDeleteGateway.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteGateway(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.gateways().deleteWithResponse("rg1", "apimService1", "gw1", "*", com.azure.core.util.Context.NONE);
    }
}
