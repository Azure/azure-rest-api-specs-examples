/** Samples for GlobalSchedules GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GlobalSchedules_Get.json
     */
    /**
     * Sample code: GlobalSchedules_Get.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void globalSchedulesGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .globalSchedules()
            .getByResourceGroupWithResponse(
                "resourceGroupName", "labvmautostart", null, com.azure.core.util.Context.NONE);
    }
}
