
/**
 * Samples for GuestSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-15/GuestSubscriptions_Get.json
     */
    /**
     * Sample code: Get a guest subscription.
     * 
     * @param manager Entry point to ComputeLimitManager.
     */
    public static void getAGuestSubscription(com.azure.resourcemanager.computelimit.ComputeLimitManager manager) {
        manager.guestSubscriptions().getWithResponse("eastus", "11111111-1111-1111-1111-111111111111",
            com.azure.core.util.Context.NONE);
    }
}
