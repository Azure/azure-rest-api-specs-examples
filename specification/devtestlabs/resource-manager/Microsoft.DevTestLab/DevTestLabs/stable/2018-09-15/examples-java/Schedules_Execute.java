
/**
 * Samples for Schedules Execute.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Schedules_Execute.json
     */
    /**
     * Sample code: Schedules_Execute.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void schedulesExecute(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.schedules().execute("resourceGroupName", "{labName}", "{scheduleName}",
            com.azure.core.util.Context.NONE);
    }
}
