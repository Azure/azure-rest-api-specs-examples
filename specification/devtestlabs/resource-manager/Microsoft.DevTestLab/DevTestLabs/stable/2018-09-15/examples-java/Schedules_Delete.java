
/**
 * Samples for Schedules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/Schedules_Delete.json
     */
    /**
     * Sample code: Schedules_Delete.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void schedulesDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.schedules().deleteWithResponse("resourceGroupName", "{labName}", "{scheduleName}",
            com.azure.core.util.Context.NONE);
    }
}
