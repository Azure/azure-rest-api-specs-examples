/** Samples for NotificationChannels List. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/NotificationChannels_List.json
     */
    /**
     * Sample code: NotificationChannels_List.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void notificationChannelsList(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .notificationChannels()
            .list("resourceGroupName", "{labName}", null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
