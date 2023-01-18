/** Samples for GlobalSchedules List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/GlobalSchedules_ListBySubscription.json
     */
    /**
     * Sample code: GlobalSchedules_ListBySubscription.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void globalSchedulesListBySubscription(
        com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.globalSchedules().list(null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
