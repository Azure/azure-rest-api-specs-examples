/** Samples for NotificationChannels Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/NotificationChannels_Delete.json
     */
    /**
     * Sample code: NotificationChannels_Delete.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void notificationChannelsDelete(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .notificationChannels()
            .deleteWithResponse(
                "resourceGroupName", "{labName}", "{notificationChannelName}", com.azure.core.util.Context.NONE);
    }
}
