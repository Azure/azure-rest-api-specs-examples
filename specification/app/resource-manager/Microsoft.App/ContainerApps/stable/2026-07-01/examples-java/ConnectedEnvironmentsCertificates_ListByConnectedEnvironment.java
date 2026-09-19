
/**
 * Samples for ConnectedEnvironmentsCertificates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ConnectedEnvironmentsCertificates_ListByConnectedEnvironment.json
     */
    /**
     * Sample code: List Certificates by Connected Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listCertificatesByConnectedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironmentsCertificates().list("examplerg", "testcontainerenv",
            com.azure.core.util.Context.NONE);
    }
}
