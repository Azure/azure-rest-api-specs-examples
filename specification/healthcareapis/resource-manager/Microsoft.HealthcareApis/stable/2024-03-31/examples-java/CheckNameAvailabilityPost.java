
import com.azure.resourcemanager.healthcareapis.models.CheckNameAvailabilityParameters;

/**
 * Samples for Services CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/
     * CheckNameAvailabilityPost.json
     */
    /**
     * Sample code: Check name availability.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void checkNameAvailability(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.services().checkNameAvailabilityWithResponse(
            new CheckNameAvailabilityParameters().withName("serviceName").withType("Microsoft.HealthcareApis/services"),
            com.azure.core.util.Context.NONE);
    }
}
