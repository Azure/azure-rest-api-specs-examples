
import com.azure.resourcemanager.postgresqlflexibleserver.models.MaintenanceEventRescheduleRequest;
import java.time.OffsetDateTime;

/**
 * Samples for MaintenanceEvents Reschedule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/MaintenanceEventsReschedule.json
     */
    /**
     * Sample code: Reschedule a maintenance event to a new date and time.
     * 
     * @param manager Entry point to PostgreSqlManager.
     */
    public static void rescheduleAMaintenanceEventToANewDateAndTime(
        com.azure.resourcemanager.postgresqlflexibleserver.PostgreSqlManager manager) {
        manager.maintenanceEvents().reschedule(
            "exampleresourcegroup", "exampleserver", "XXXX-111", new MaintenanceEventRescheduleRequest()
                .withPostponeToDateTime(OffsetDateTime.parse("2026-04-10T10:00:00+00:00")),
            com.azure.core.util.Context.NONE);
    }
}
