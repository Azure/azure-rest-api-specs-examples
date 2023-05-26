/** Samples for ImpactedResources ListBySubscriptionIdAndEventId. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/ImpactedResources_ListBySubscriptionId_ListByEventId.json
     */
    /**
     * Sample code: ListImpactedResourcesBySubscriptionId.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void listImpactedResourcesBySubscriptionId(
        com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager
            .impactedResources()
            .listBySubscriptionIdAndEventId("BC_1-FXZ", "targetRegion eq 'westus'", com.azure.core.util.Context.NONE);
    }
}
