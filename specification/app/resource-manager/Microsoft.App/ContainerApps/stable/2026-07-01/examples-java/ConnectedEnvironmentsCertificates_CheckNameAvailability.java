
import com.azure.resourcemanager.appcontainers.models.CheckNameAvailabilityRequest;

/**
 * Samples for ConnectedEnvironments CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ConnectedEnvironmentsCertificates_CheckNameAvailability.json
     */
    /**
     * Sample code: Certificates_CheckNameAvailability.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        certificatesCheckNameAvailability(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironments()
            .checkNameAvailabilityWithResponse("examplerg", "testcontainerenv", new CheckNameAvailabilityRequest()
                .withName("testcertificatename").withType("Microsoft.App/connectedEnvironments/certificates"),
                com.azure.core.util.Context.NONE);
    }
}
