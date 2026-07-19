
import com.azure.resourcemanager.sql.models.StartStopScheduleName;

/**
 * Samples for StartStopManagedInstanceSchedules Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/StartStopManagedInstanceScheduleGet.json
     */
    /**
     * Sample code: Gets the managed instance's Start/Stop schedule.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsTheManagedInstanceSStartStopSchedule(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getStartStopManagedInstanceSchedules().getWithResponse("schedulerg", "schedulemi",
            StartStopScheduleName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
