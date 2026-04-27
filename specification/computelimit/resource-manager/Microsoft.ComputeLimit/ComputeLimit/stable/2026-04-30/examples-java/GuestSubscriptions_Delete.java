
/**
 * Samples for GuestSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/GuestSubscriptions_Delete.json
     */
    /**
     * Sample code: Delete a guest subscription.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void deleteAGuestSubscription(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.guestSubscriptions().deleteByResourceGroupWithResponse("eastus", "11111111-1111-1111-1111-111111111111",
            com.azure.core.util.Context.NONE);
    }
}
