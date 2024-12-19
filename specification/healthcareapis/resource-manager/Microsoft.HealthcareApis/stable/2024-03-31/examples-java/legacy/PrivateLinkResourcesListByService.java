
/**
 * Samples for PrivateLinkResources ListByService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/
     * PrivateLinkResourcesListByService.json
     */
    /**
     * Sample code: PrivateLinkResources_ListGroupIds.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void
        privateLinkResourcesListGroupIds(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.privateLinkResources().listByServiceWithResponse("rgname", "service1",
            com.azure.core.util.Context.NONE);
    }
}
