
/**
 * Samples for Services GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/
     * ServiceGet.json
     */
    /**
     * Sample code: Get metadata.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void getMetadata(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.services().getByResourceGroupWithResponse("rg1", "service1", com.azure.core.util.Context.NONE);
    }
}
