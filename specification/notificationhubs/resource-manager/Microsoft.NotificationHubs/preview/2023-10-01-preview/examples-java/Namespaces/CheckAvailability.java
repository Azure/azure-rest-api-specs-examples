
import com.azure.resourcemanager.notificationhubs.models.CheckAvailabilityParameters;

/**
 * Samples for Namespaces CheckAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/preview/2023-10-01-preview/examples/
     * Namespaces/CheckAvailability.json
     */
    /**
     * Sample code: Namespaces_CheckAvailability.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        namespacesCheckAvailability(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().checkAvailabilityWithResponse(
            new CheckAvailabilityParameters().withName("sdk-Namespace-2924"), com.azure.core.util.Context.NONE);
    }
}
