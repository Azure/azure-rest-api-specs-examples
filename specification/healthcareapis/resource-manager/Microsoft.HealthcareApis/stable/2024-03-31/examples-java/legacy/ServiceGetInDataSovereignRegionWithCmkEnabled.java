
/**
 * Samples for Services GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/
     * ServiceGetInDataSovereignRegionWithCmkEnabled.json
     */
    /**
     * Sample code: Get metadata for CMK enabled service in data sovereign region.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void getMetadataForCMKEnabledServiceInDataSovereignRegion(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.services().getByResourceGroupWithResponse("rg1", "service1", com.azure.core.util.Context.NONE);
    }
}
