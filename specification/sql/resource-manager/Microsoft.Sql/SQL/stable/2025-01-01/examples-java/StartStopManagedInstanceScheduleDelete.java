
import com.azure.resourcemanager.sql.models.StartStopScheduleName;

/**
 * Samples for StartStopManagedInstanceSchedules Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/StartStopManagedInstanceScheduleDelete.json
     */
    /**
     * Sample code: Deletes the managed instance's Start/Stop schedule.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        deletesTheManagedInstanceSStartStopSchedule(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getStartStopManagedInstanceSchedules().deleteWithResponse("schedulerg", "schedulemi",
            StartStopScheduleName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
