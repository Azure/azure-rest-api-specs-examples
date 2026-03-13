
/**
 * Samples for Schedulers GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/Schedulers_Get.json
     */
    /**
     * Sample code: Schedulers_Get.
     * 
     * @param manager Entry point to DurableTaskManager.
     */
    public static void schedulersGet(com.azure.resourcemanager.durabletask.DurableTaskManager manager) {
        manager.schedulers().getByResourceGroupWithResponse("rgopenapi", "testscheduler",
            com.azure.core.util.Context.NONE);
    }
}
