```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.labservices.models.RecurrenceFrequency;
import com.azure.resourcemanager.labservices.models.RecurrencePattern;
import com.azure.resourcemanager.labservices.models.Schedule;
import java.time.LocalDate;

/** Samples for Schedules Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Schedules/patchSchedule.json
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
                    .withExpirationDate(LocalDate.parse("2020-08-14")))
            .apply();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-labservices_1.0.0-beta.2/sdk/labservices/azure-resourcemanager-labservices/README.md) on how to add the SDK to your project and authenticate.
