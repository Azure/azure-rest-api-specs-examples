/** Samples for AvailabilityStatuses ListBySubscriptionId. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/AvailabilityStatuses_ListBySubscriptionId.json
     */
    /**
     * Sample code: ListHealthBySubscriptionId.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void listHealthBySubscriptionId(
        com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager
            .availabilityStatuses()
            .listBySubscriptionId(null, "recommendedactions", com.azure.core.util.Context.NONE);
    }
}
