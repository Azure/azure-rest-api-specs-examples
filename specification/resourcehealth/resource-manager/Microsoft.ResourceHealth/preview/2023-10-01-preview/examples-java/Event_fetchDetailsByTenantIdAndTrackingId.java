/** Samples for EventOperation FetchDetailsByTenantIdAndTrackingId. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/Event_fetchDetailsByTenantIdAndTrackingId.json
     */
    /**
     * Sample code: EventDetailsByTenantIdAndTrackingId.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void eventDetailsByTenantIdAndTrackingId(
        com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager
            .eventOperations()
            .fetchDetailsByTenantIdAndTrackingIdWithResponse("eventTrackingId", com.azure.core.util.Context.NONE);
    }
}
