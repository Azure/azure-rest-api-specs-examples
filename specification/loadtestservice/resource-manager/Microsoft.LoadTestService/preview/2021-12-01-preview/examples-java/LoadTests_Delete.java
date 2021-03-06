import com.azure.core.util.Context;

/** Samples for LoadTests Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/preview/2021-12-01-preview/examples/LoadTests_Delete.json
     */
    /**
     * Sample code: LoadTests_Delete.
     *
     * @param manager Entry point to LoadTestManager.
     */
    public static void loadTestsDelete(com.azure.resourcemanager.loadtestservice.LoadTestManager manager) {
        manager.loadTests().delete("dummyrg", "myLoadTest", Context.NONE);
    }
}
