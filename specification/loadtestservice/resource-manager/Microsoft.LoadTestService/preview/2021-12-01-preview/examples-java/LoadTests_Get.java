import com.azure.core.util.Context;

/** Samples for LoadTests GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/preview/2021-12-01-preview/examples/LoadTests_Get.json
     */
    /**
     * Sample code: LoadTests_Get.
     *
     * @param manager Entry point to LoadTestManager.
     */
    public static void loadTestsGet(com.azure.resourcemanager.loadtestservice.LoadTestManager manager) {
        manager.loadTests().getByResourceGroupWithResponse("dummyrg", "myLoadTest", Context.NONE);
    }
}
