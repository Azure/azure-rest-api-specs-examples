
import com.azure.resourcemanager.devtestlabs.fluent.models.ScheduleInner;
import com.azure.resourcemanager.devtestlabs.models.DayDetails;
import com.azure.resourcemanager.devtestlabs.models.EnableStatus;
import com.azure.resourcemanager.devtestlabs.models.HourDetails;
import com.azure.resourcemanager.devtestlabs.models.NotificationSettings;
import com.azure.resourcemanager.devtestlabs.models.WeekDetails;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Schedules CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/
     * Schedules_CreateOrUpdate.json
     */
    /**
     * Sample code: Schedules_CreateOrUpdate.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void schedulesCreateOrUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.schedules().createOrUpdateWithResponse("resourceGroupName", "{labName}", "{scheduleName}",
            new ScheduleInner().withLocation("{location}").withTags(mapOf("tagName1", "tagValue1"))
                .withStatus(EnableStatus.fromString("{Enabled|Disabled}")).withTaskType("{myLabVmTaskType}")
                .withWeeklyRecurrence(new WeekDetails().withWeekdays(Arrays.asList("Monday", "Wednesday", "Friday"))
                    .withTime("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"))
                .withDailyRecurrence(new DayDetails().withTime("{timeOfTheDayTheScheduleWillOccurEveryDay}"))
                .withHourlyRecurrence(new HourDetails().withMinute(30)).withTimeZoneId("Pacific Standard Time")
                .withNotificationSettings(new NotificationSettings()
                    .withStatus(EnableStatus.fromString("{Enabled|Disabled}")).withTimeInMinutes(15)
                    .withWebhookUrl("{webhookUrl}").withEmailRecipient("{email}").withNotificationLocale("EN"))
                .withTargetResourceId(
                    "/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
