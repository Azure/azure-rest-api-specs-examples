
/**
 * Samples for Quotas CheckAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/
     * Quotas_CheckAvailability.json
     */
    /**
     * Sample code: Check Quota Availability on quota bucket per region per subscription.
     * 
     * @param manager Entry point to LoadTestManager.
     */
    public static void checkQuotaAvailabilityOnQuotaBucketPerRegionPerSubscription(
        com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        manager.quotas().checkAvailabilityWithResponse("westus", "testQuotaBucket", null,
            com.azure.core.util.Context.NONE);
    }
}
