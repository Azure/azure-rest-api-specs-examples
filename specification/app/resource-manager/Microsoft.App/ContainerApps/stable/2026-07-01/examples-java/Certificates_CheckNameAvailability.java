
import com.azure.resourcemanager.appcontainers.models.CheckNameAvailabilityRequest;

/**
 * Samples for Namespaces CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Certificates_CheckNameAvailability.json
     */
    /**
     * Sample code: Certificates_CheckNameAvailability.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        certificatesCheckNameAvailability(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.namespaces()
            .checkNameAvailabilityWithResponse("examplerg", "testcontainerenv", new CheckNameAvailabilityRequest()
                .withName("testcertificatename").withType("Microsoft.App/managedEnvironments/certificates"),
                com.azure.core.util.Context.NONE);
    }
}
