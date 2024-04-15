
/**
 * Samples for Services Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/
     * ServiceDelete.json
     */
    /**
     * Sample code: Delete service.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void deleteService(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.services().delete("rg1", "service1", com.azure.core.util.Context.NONE);
    }
}
