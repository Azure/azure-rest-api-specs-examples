
/**
 * Samples for Quotas Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/Quotas_Get.
     * json
     */
    /**
     * Sample code: Get the available quota for a quota bucket per region per subscription.
     * 
     * @param manager Entry point to LoadTestManager.
     */
    public static void getTheAvailableQuotaForAQuotaBucketPerRegionPerSubscription(
        com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        manager.quotas().getWithResponse("westus", "testQuotaBucket", com.azure.core.util.Context.NONE);
    }
}
