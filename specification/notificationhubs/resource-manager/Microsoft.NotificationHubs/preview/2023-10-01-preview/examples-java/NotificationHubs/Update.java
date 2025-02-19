
import com.azure.resourcemanager.notificationhubs.models.GcmCredential;
import com.azure.resourcemanager.notificationhubs.models.NotificationHubResource;

/**
 * Samples for NotificationHubs Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * NotificationHubs/Update.json
     */
    /**
     * Sample code: NotificationHubs_Update.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        notificationHubsUpdate(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        NotificationHubResource resource = manager.notificationHubs().getWithResponse("sdkresourceGroup", "nh-sdk-ns",
            "sdk-notificationHubs-8708", com.azure.core.util.Context.NONE).getValue();
        resource
            .update().withRegistrationTtl("10675199.02:48:05.4775807").withGcmCredential(new GcmCredential()
                .withGcmEndpoint("https://fcm.googleapis.com/fcm/send").withGoogleApiKey("fakeTokenPlaceholder"))
            .apply();
    }
}
