import com.azure.core.util.Context;
import com.azure.resourcemanager.labservices.models.RecurrenceFrequency;
import com.azure.resourcemanager.labservices.models.RecurrencePattern;
import com.azure.resourcemanager.labservices.models.Schedule;
import java.time.OffsetDateTime;

/** Samples for Schedules Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Schedules/patchSchedule.json
     */
    /**
     * Sample code: patchSchedule.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void patchSchedule(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        Schedule resource =
            manager.schedules().getWithResponse("testrg123", "testlab", "schedule1", Context.NONE).getValue();
        resource
            .update()
            .withRecurrencePattern(
                new RecurrencePattern()
                    .withFrequency(RecurrenceFrequency.DAILY)
                    .withInterval(2)
                    .withExpirationDate(OffsetDateTime.parse("2020-08-14T23:59:59Z")))
            .apply();
    }
}
