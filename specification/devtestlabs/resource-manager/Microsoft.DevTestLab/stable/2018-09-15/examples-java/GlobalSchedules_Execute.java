/** Samples for GlobalSchedules Execute. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GlobalSchedules_Execute.json
     */
    /**
     * Sample code: GlobalSchedules_Execute.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void globalSchedulesExecute(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.globalSchedules().execute("resourceGroupName", "labvmautostart", com.azure.core.util.Context.NONE);
    }
}
