/** Samples for NetworkStatus ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementServiceGetNetworkStatusByLocation.json
     */
    /**
     * Sample code: ApiManagementServiceGetNetworkStatusByLocation.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementServiceGetNetworkStatusByLocation(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .networkStatus()
            .listByLocationWithResponse("rg1", "apimService1", "North Central US", com.azure.core.util.Context.NONE);
    }
}
