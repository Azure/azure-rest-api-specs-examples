
/**
 * Samples for StartStopManagedInstanceSchedules ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/StartStopManagedInstanceScheduleList.json
     */
    /**
     * Sample code: Lists the managed instance's Start/Stop schedules.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listsTheManagedInstanceSStartStopSchedules(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getStartStopManagedInstanceSchedules().listByInstance("schedulerg", "schedulemi",
            com.azure.core.util.Context.NONE);
    }
}
