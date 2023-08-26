/** Samples for QuotaByCounterKeys ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetQuotaCounterKeys.json
     */
    /**
     * Sample code: ApiManagementGetQuotaCounterKeys.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetQuotaCounterKeys(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .quotaByCounterKeys()
            .listByServiceWithResponse("rg1", "apimService1", "ba", com.azure.core.util.Context.NONE);
    }
}
