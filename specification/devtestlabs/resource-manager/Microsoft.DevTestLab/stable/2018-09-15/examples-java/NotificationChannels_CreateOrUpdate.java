import com.azure.resourcemanager.devtestlabs.models.Event;
import com.azure.resourcemanager.devtestlabs.models.NotificationChannelEventType;
import java.util.Arrays;

/** Samples for NotificationChannels CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/NotificationChannels_CreateOrUpdate.json
     */
    /**
     * Sample code: NotificationChannels_CreateOrUpdate.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void notificationChannelsCreateOrUpdate(
        com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .notificationChannels()
            .define("{notificationChannelName}")
            .withRegion((String) null)
            .withExistingLab("resourceGroupName", "{labName}")
            .withWebhookUrl("{webhookUrl}")
            .withEmailRecipient("{email}")
            .withNotificationLocale("en")
            .withDescription("Integration configured for auto-shutdown")
            .withEvents(Arrays.asList(new Event().withEventName(NotificationChannelEventType.AUTO_SHUTDOWN)))
            .create();
    }
}
