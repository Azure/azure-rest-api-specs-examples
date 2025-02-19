
/**
 * Samples for EventOperation FetchDetailsBySubscriptionIdAndTrackingId.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/
     * Event_fetchDetailsBySubscriptionIdAndTrackingId.json
     */
    /**
     * Sample code: EventDetailsBySubscriptionIdAndTrackingId.
     * 
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void eventDetailsBySubscriptionIdAndTrackingId(
        com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager.eventOperations().fetchDetailsBySubscriptionIdAndTrackingIdWithResponse("eventTrackingId",
            com.azure.core.util.Context.NONE);
    }
}
