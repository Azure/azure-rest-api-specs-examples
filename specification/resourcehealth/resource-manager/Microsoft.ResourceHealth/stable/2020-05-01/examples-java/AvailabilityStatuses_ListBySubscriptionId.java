import com.azure.core.util.Context;

/** Samples for AvailabilityStatuses ListBySubscriptionId. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2020-05-01/examples/AvailabilityStatuses_ListBySubscriptionId.json
     */
    /**
     * Sample code: ListHealthBySubscriptionId.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void listHealthBySubscriptionId(
        com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager.availabilityStatuses().listBySubscriptionId(null, "recommendedactions", Context.NONE);
    }
}
