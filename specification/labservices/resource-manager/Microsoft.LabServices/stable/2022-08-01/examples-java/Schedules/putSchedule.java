
import com.azure.resourcemanager.labservices.models.RecurrenceFrequency;
import com.azure.resourcemanager.labservices.models.RecurrencePattern;
import java.time.OffsetDateTime;

/**
 * Samples for Schedules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/Schedules/putSchedule
     * .json
     */
    /**
     * Sample code: putSchedule.
     * 
     * @param manager Entry point to LabServicesManager.
     */
    public static void putSchedule(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.schedules().define("schedule1").withExistingLab("testrg123", "testlab")
            .withStartAt(OffsetDateTime.parse("2020-05-26T12:00:00Z"))
            .withStopAt(OffsetDateTime.parse("2020-05-26T18:00:00Z"))
            .withRecurrencePattern(new RecurrencePattern().withFrequency(RecurrenceFrequency.DAILY).withInterval(2)
                .withExpirationDate(OffsetDateTime.parse("2020-08-14T23:59:59Z")))
            .withTimeZoneId("America/Los_Angeles").withNotes("Schedule 1 for students").create();
    }
}
