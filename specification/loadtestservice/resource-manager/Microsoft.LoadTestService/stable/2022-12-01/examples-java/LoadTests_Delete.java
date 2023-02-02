/** Samples for LoadTests Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/LoadTests_Delete.json
     */
    /**
     * Sample code: LoadTests_Delete.
     *
     * @param manager Entry point to LoadTestManager.
     */
    public static void loadTestsDelete(com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        manager.loadTests().delete("dummyrg", "myLoadTest", com.azure.core.util.Context.NONE);
    }
}
