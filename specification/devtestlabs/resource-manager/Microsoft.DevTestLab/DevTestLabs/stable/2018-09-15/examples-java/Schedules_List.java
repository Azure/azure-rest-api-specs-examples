
/**
 * Samples for Schedules List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Schedules_List.json
     */
    /**
     * Sample code: Schedules_List.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void schedulesList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.schedules().list("resourceGroupName", "{labName}", null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
