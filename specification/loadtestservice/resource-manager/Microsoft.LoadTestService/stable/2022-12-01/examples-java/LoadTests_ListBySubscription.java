import com.azure.core.util.Context;

/** Samples for LoadTests List. */
public final class Main {
    /*
     * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_ListBySubscription.json
     */
    /**
     * Sample code: LoadTests_ListBySubscription.
     *
     * @param manager Entry point to LoadTestManager.
     */
    public static void loadTestsListBySubscription(com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        manager.loadTests().list(Context.NONE);
    }
}
