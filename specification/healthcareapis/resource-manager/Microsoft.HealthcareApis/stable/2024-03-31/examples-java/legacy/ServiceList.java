
/**
 * Samples for Services List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/legacy/
     * ServiceList.json
     */
    /**
     * Sample code: List all services in subscription.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void
        listAllServicesInSubscription(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.services().list(com.azure.core.util.Context.NONE);
    }
}
