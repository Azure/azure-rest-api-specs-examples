
/**
 * Samples for LoadTests List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/
     * LoadTests_ListBySubscription.json
     */
    /**
     * Sample code: List LoadTestResource resources by subscription ID.
     * 
     * @param manager Entry point to LoadTestManager.
     */
    public static void
        listLoadTestResourceResourcesBySubscriptionID(com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        manager.loadTests().list(com.azure.core.util.Context.NONE);
    }
}
