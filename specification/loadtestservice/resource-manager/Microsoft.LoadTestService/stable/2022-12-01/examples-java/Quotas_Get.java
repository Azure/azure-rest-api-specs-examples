import com.azure.core.util.Context;

/** Samples for Quotas Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/Quotas_Get.json
     */
    /**
     * Sample code: Quotas_Get.
     *
     * @param manager Entry point to LoadTestManager.
     */
    public static void quotasGet(com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        manager.quotas().getWithResponse("westus", "testQuotaBucket", Context.NONE);
    }
}
