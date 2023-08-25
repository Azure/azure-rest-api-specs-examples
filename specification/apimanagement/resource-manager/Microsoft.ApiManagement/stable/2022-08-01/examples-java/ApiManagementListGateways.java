/** Samples for Gateway ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListGateways.json
     */
    /**
     * Sample code: ApiManagementListGateways.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListGateways(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.gateways().listByService("rg1", "apimService1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
