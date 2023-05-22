import com.azure.resourcemanager.costmanagement.models.DaysOfWeek;
import com.azure.resourcemanager.costmanagement.models.FileDestination;
import com.azure.resourcemanager.costmanagement.models.FileFormat;
import com.azure.resourcemanager.costmanagement.models.NotificationProperties;
import com.azure.resourcemanager.costmanagement.models.ScheduleFrequency;
import com.azure.resourcemanager.costmanagement.models.ScheduleProperties;
import com.azure.resourcemanager.costmanagement.models.ScheduledActionKind;
import com.azure.resourcemanager.costmanagement.models.ScheduledActionStatus;
import com.azure.resourcemanager.costmanagement.models.WeeksOfMonth;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for ScheduledActions CreateOrUpdateByScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledAction-createOrUpdate-shared.json
     */
    /**
     * Sample code: CreateOrUpdateScheduledActionByScope.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void createOrUpdateScheduledActionByScope(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .scheduledActions()
            .define("monthlyCostByResource")
            .withExistingScope("subscriptions/00000000-0000-0000-0000-000000000000")
            .withKind(ScheduledActionKind.EMAIL)
            .withDisplayName("Monthly Cost By Resource")
            .withFileDestination(new FileDestination().withFileFormats(Arrays.asList(FileFormat.CSV)))
            .withNotification(
                new NotificationProperties()
                    .withTo(Arrays.asList("user@gmail.com", "team@gmail.com"))
                    .withSubject("Cost by resource this month"))
            .withSchedule(
                new ScheduleProperties()
                    .withFrequency(ScheduleFrequency.MONTHLY)
                    .withHourOfDay(10)
                    .withDaysOfWeek(Arrays.asList(DaysOfWeek.MONDAY))
                    .withWeeksOfMonth(Arrays.asList(WeeksOfMonth.FIRST, WeeksOfMonth.THIRD))
                    .withStartDate(OffsetDateTime.parse("2020-06-19T22:21:51.1287144Z"))
                    .withEndDate(OffsetDateTime.parse("2021-06-19T22:21:51.1287144Z")))
            .withStatus(ScheduledActionStatus.ENABLED)
            .withViewId("/providers/Microsoft.CostManagement/views/swaggerExample")
            .withIfMatch("")
            .create();
    }
}
