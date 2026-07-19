
import com.azure.resourcemanager.sql.fluent.models.StartStopManagedInstanceScheduleInner;
import com.azure.resourcemanager.sql.models.DayOfWeek;
import com.azure.resourcemanager.sql.models.ScheduleItem;
import com.azure.resourcemanager.sql.models.StartStopScheduleName;
import java.util.Arrays;

/**
 * Samples for StartStopManagedInstanceSchedules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/StartStopManagedInstanceScheduleCreateOrUpdateMin.json
     */
    /**
     * Sample code: Creates or updates the managed instance's Start/Stop schedule with no optional parameters specified.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createsOrUpdatesTheManagedInstanceSStartStopScheduleWithNoOptionalParametersSpecified(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getStartStopManagedInstanceSchedules().createOrUpdateWithResponse("schedulerg",
            "schedulemi", StartStopScheduleName.DEFAULT,
            new StartStopManagedInstanceScheduleInner().withScheduleList(Arrays.asList(
                new ScheduleItem().withStartDay(DayOfWeek.THURSDAY).withStartTime("18:00")
                    .withStopDay(DayOfWeek.THURSDAY).withStopTime("17:00"),
                new ScheduleItem().withStartDay(DayOfWeek.THURSDAY).withStartTime("15:00")
                    .withStopDay(DayOfWeek.THURSDAY).withStopTime("14:00"))),
            com.azure.core.util.Context.NONE);
    }
}
