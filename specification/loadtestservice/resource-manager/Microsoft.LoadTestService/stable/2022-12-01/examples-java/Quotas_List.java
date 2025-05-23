
/**
 * Samples for Quotas List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/Quotas_List.
     * json
     */
    /**
     * Sample code: List quotas for a given subscription Id.
     * 
     * @param manager Entry point to LoadTestManager.
     */
    public static void
        listQuotasForAGivenSubscriptionId(com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        manager.quotas().list("westus", com.azure.core.util.Context.NONE);
    }
}
