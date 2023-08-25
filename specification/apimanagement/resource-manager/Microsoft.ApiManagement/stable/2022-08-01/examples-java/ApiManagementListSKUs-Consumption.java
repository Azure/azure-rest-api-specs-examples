/** Samples for ApiManagementServiceSkus ListAvailableServiceSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListSKUs-Consumption.json
     */
    /**
     * Sample code: ApiManagementListSKUs-Consumption.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListSKUsConsumption(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiManagementServiceSkus()
            .listAvailableServiceSkus("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
