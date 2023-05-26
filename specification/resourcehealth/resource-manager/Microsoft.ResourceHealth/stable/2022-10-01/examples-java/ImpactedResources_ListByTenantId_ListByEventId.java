/** Samples for ImpactedResources ListByTenantIdAndEventId. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/ImpactedResources_ListByTenantId_ListByEventId.json
     */
    /**
     * Sample code: ListEventsByTenantId.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void listEventsByTenantId(com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager
            .impactedResources()
            .listByTenantIdAndEventId("BC_1-FXZ", "targetRegion eq 'westus'", com.azure.core.util.Context.NONE);
    }
}
