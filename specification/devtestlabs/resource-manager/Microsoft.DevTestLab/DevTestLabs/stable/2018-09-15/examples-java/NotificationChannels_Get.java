
/**
 * Samples for NotificationChannels Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-09-15/NotificationChannels_Get.json
     */
    /**
     * Sample code: NotificationChannels_Get.
     * 
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void notificationChannelsGet(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager.notificationChannels().getWithResponse("resourceGroupName", "{labName}", "{notificationChannelName}",
            null, com.azure.core.util.Context.NONE);
    }
}
