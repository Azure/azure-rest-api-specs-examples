
import com.azure.resourcemanager.notificationhubs.models.CheckAvailabilityParameters;

/**
 * Samples for Namespaces CheckAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/notificationhubs/resource-manager/Microsoft.NotificationHubs/stable/2017-04-01/examples/Namespaces/
     * NHNameSpaceCheckNameAvailability.json
     */
    /**
     * Sample code: NameSpaceCheckNameAvailability.
     * 
     * @param manager Entry point to NotificationHubsManager.
     */
    public static void
        nameSpaceCheckNameAvailability(com.azure.resourcemanager.notificationhubs.NotificationHubsManager manager) {
        manager.namespaces().checkAvailabilityWithResponse(
            new CheckAvailabilityParameters().withName("sdk-Namespace-2924"), com.azure.core.util.Context.NONE);
    }
}
