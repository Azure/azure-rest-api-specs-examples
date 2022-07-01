import com.azure.core.util.Context;

/** Samples for WordpressInstances Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/phpworkloads/WordpressInstances_Get.json
     */
    /**
     * Sample code: Workloads.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void workloads(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.wordpressInstances().getWithResponse("test-rg", "wp39", Context.NONE);
    }
}
