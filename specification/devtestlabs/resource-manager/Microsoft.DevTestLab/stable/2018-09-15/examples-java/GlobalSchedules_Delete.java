/** Samples for GlobalSchedules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GlobalSchedules_Delete.json
     */
    /**
     * Sample code: GlobalSchedules_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void globalSchedulesDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .globalSchedules()
            .deleteByResourceGroupWithResponse("resourceGroupName", "labvmautostart", com.azure.core.util.Context.NONE);
    }
}
