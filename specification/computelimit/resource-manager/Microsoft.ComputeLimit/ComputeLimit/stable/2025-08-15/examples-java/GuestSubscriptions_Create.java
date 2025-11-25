
import com.azure.resourcemanager.computelimit.models.GuestSubscriptionProperties;

/**
 * Samples for GuestSubscriptions Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-15/GuestSubscriptions_Create.json
     */
    /**
     * Sample code: Create a guest subscription.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void createAGuestSubscription(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.guestSubscriptions().define("11111111-1111-1111-1111-111111111111").withExistingLocation("eastus")
            .withProperties(new GuestSubscriptionProperties()).create();
    }
}
