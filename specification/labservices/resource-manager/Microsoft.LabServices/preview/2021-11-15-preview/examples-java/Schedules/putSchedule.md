Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-labservices_1.0.0-beta.2/sdk/labservices/azure-resourcemanager-labservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.labservices.models.RecurrenceFrequency;
import com.azure.resourcemanager.labservices.models.RecurrencePattern;
import java.time.LocalDate;
import java.time.OffsetDateTime;

/** Samples for Schedules CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Schedules/putSchedule.json
     */
    /**
     * Sample code: putSchedule.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void putSchedule(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager
            .schedules()
            .define("schedule1")
            .withExistingLab("testrg123", "testlab")
            .withStartAt(OffsetDateTime.parse("2020-05-26T12:00:00Z"))
            .withStopAt(OffsetDateTime.parse("2020-05-26T18:00:00Z"))
            .withRecurrencePattern(
                new RecurrencePattern()
                    .withFrequency(RecurrenceFrequency.DAILY)
                    .withInterval(2)
                    .withExpirationDate(LocalDate.parse("2020-08-14")))
            .withTimeZoneId("America/Los_Angeles")
            .withNotes("Schedule 1 for students")
            .create();
    }
}
```
