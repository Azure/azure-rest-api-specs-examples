
import com.azure.resourcemanager.devtestlabs.models.NotificationChannelEventType;
import com.azure.resourcemanager.devtestlabs.models.NotifyParameters;

/**
 * Samples for NotificationChannels Notify.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/NotificationChannels_Notify.json
     */
    /**
     * Sample code: NotificationChannels_Notify.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void notificationChannelsNotify(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.notificationChannels().notifyWithResponse("resourceGroupName", "{labName}", "{notificationChannelName}",
            new NotifyParameters().withEventName(NotificationChannelEventType.AUTO_SHUTDOWN).withJsonPayload(
                "{\"eventType\":\"AutoShutdown\",\"subscriptionId\":\"{subscriptionId}\",\"resourceGroupName\":\"resourceGroupName\",\"labName\":\"{labName}\"}"),
            com.azure.core.util.Context.NONE);
    }
}
