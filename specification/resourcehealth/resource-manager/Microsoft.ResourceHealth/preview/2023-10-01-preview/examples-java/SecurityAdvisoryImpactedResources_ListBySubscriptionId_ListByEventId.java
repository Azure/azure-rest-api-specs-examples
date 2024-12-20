
/**
 * Samples for SecurityAdvisoryImpactedResources ListBySubscriptionIdAndEventId.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/
     * SecurityAdvisoryImpactedResources_ListBySubscriptionId_ListByEventId.json
     */
    /**
     * Sample code: ListSecurityAdvisoryImpactedResourcesBySubscriptionId.
     * 
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void listSecurityAdvisoryImpactedResourcesBySubscriptionId(
        com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager.securityAdvisoryImpactedResources().listBySubscriptionIdAndEventId("BC_1-FXZ", null,
            com.azure.core.util.Context.NONE);
    }
}
