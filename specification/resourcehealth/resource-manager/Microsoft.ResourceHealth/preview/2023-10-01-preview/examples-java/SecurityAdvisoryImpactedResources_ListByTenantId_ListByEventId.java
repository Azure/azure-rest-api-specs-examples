/** Samples for SecurityAdvisoryImpactedResources ListByTenantIdAndEventId. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/SecurityAdvisoryImpactedResources_ListByTenantId_ListByEventId.json
     */
    /**
     * Sample code: ListSecurityAdvisoryImpactedResourcesByTenantId.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void listSecurityAdvisoryImpactedResourcesByTenantId(
        com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager
            .securityAdvisoryImpactedResources()
            .listByTenantIdAndEventId("BC_1-FXZ", null, com.azure.core.util.Context.NONE);
    }
}
