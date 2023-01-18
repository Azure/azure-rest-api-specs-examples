/** Samples for Schedules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Schedules_List.json
     */
    /**
     * Sample code: Schedules_List.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void schedulesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .schedules()
            .list("resourceGroupName", "{labName}", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
