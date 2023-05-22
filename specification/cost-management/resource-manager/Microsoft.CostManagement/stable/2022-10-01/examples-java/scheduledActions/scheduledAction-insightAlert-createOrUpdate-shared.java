import com.azure.resourcemanager.costmanagement.models.NotificationProperties;
import com.azure.resourcemanager.costmanagement.models.ScheduleFrequency;
import com.azure.resourcemanager.costmanagement.models.ScheduleProperties;
import com.azure.resourcemanager.costmanagement.models.ScheduledActionKind;
import com.azure.resourcemanager.costmanagement.models.ScheduledActionStatus;
import java.time.OffsetDateTime;
import java.util.Arrays;

/** Samples for ScheduledActions CreateOrUpdateByScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/cost-management/resource-manager/Microsoft.CostManagement/stable/2022-10-01/examples/scheduledActions/scheduledAction-insightAlert-createOrUpdate-shared.json
     */
    /**
     * Sample code: CreateOrUpdateInsightAlertScheduledActionByScope.
     *
     * @param manager Entry point to CostManagementManager.
     */
    public static void createOrUpdateInsightAlertScheduledActionByScope(
        com.azure.resourcemanager.costmanagement.CostManagementManager manager) {
        manager
            .scheduledActions()
            .define("dailyAnomalyByResource")
            .withExistingScope("subscriptions/00000000-0000-0000-0000-000000000000")
            .withKind(ScheduledActionKind.INSIGHT_ALERT)
            .withDisplayName("Daily anomaly by resource")
            .withNotification(
                new NotificationProperties()
                    .withTo(Arrays.asList("user@gmail.com", "team@gmail.com"))
                    .withSubject("Cost anomaly detected in the resource"))
            .withSchedule(
                new ScheduleProperties()
                    .withFrequency(ScheduleFrequency.DAILY)
                    .withStartDate(OffsetDateTime.parse("2020-06-19T22:21:51.1287144Z"))
                    .withEndDate(OffsetDateTime.parse("2021-06-19T22:21:51.1287144Z")))
            .withStatus(ScheduledActionStatus.ENABLED)
            .withViewId("/providers/Microsoft.CostManagement/views/swaggerExample")
            .withIfMatch("")
            .create();
    }
}
