/** Samples for EventOperation GetBySubscriptionIdAndTrackingId. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/Event_GetBySubscriptionIdAndTrackingId.json
     */
    /**
     * Sample code: SecurityAdvisoriesEventBySubscriptionIdAndTrackingId.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void securityAdvisoriesEventBySubscriptionIdAndTrackingId(
        com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager
            .eventOperations()
            .getBySubscriptionIdAndTrackingIdWithResponse(
                "eventTrackingId", "properties/status eq 'Active'", "7/10/2022", com.azure.core.util.Context.NONE);
    }
}
