/** Samples for QuotaByPeriodKeys Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetQuotaCounterKeysByQuotaPeriod.json
     */
    /**
     * Sample code: ApiManagementGetQuotaCounterKeysByQuotaPeriod.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetQuotaCounterKeysByQuotaPeriod(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .quotaByPeriodKeys()
            .getWithResponse("rg1", "apimService1", "ba", "0_P3Y6M4DT12H30M5S", com.azure.core.util.Context.NONE);
    }
}
