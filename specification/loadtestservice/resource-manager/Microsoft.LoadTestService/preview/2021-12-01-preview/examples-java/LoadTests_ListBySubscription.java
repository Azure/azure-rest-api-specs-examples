import com.azure.core.util.Context;

/** Samples for LoadTests List. */
public final class Main {
    /*
     * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/preview/2021-12-01-preview/examples/LoadTests_ListBySubscription.json
     */
    /**
     * Sample code: LoadTests_ListBySubscription.
     *
     * @param manager Entry point to LoadTestManager.
     */
    public static void loadTestsListBySubscription(com.azure.resourcemanager.loadtestservice.LoadTestManager manager) {
        manager.loadTests().list(Context.NONE);
    }
}
