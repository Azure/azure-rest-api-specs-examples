/** Samples for Schedules ListApplicable. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Schedules_ListApplicable.json
     */
    /**
     * Sample code: Schedules_ListApplicable.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void schedulesListApplicable(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .schedules()
            .listApplicable("resourceGroupName", "{labName}", "{scheduleName}", com.azure.core.util.Context.NONE);
    }
}
