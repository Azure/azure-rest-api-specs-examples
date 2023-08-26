/** Samples for NetworkStatus ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceGetNetworkStatus.json
     */
    /**
     * Sample code: ApiManagementServiceGetNetworkStatus.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementServiceGetNetworkStatus(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.networkStatus().listByServiceWithResponse("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
