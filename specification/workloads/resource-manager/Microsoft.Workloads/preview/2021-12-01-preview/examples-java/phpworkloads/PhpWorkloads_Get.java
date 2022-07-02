import com.azure.core.util.Context;

/** Samples for PhpWorkloads GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/phpworkloads/PhpWorkloads_Get.json
     */
    /**
     * Sample code: Workloads.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void workloads(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.phpWorkloads().getByResourceGroupWithResponse("test-rg", "wp39", Context.NONE);
    }
}
