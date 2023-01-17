/** Samples for Schedules Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Schedules_Get.json
     */
    /**
     * Sample code: Schedules_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void schedulesGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .schedules()
            .getWithResponse(
                "resourceGroupName", "{labName}", "{scheduleName}", null, com.azure.core.util.Context.NONE);
    }
}
