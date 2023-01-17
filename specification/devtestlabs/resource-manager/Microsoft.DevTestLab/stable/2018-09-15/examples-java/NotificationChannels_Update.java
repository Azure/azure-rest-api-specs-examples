import com.azure.resourcemanager.devtestlabs.models.NotificationChannel;

/** Samples for NotificationChannels Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/NotificationChannels_Update.json
     */
    /**
     * Sample code: NotificationChannels_Update.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void notificationChannelsUpdate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        NotificationChannel resource =
            manager
                .notificationChannels()
                .getWithResponse(
                    "resourceGroupName",
                    "{labName}",
                    "{notificationChannelName}",
                    null,
                    com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
