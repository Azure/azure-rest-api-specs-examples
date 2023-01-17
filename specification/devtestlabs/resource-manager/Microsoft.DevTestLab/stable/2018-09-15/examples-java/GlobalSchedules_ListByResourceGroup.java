/** Samples for GlobalSchedules ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GlobalSchedules_ListByResourceGroup.json
     */
    /**
     * Sample code: GlobalSchedules_ListByResourceGroup.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void globalSchedulesListByResourceGroup(
        com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .globalSchedules()
            .listByResourceGroup("resourceGroupName", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
