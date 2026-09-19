
/**
 * Samples for ConnectedEnvironmentsCertificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ConnectedEnvironmentsCertificate_Get.json
     */
    /**
     * Sample code: Get Certificate.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getCertificate(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.connectedEnvironmentsCertificates().getWithResponse("examplerg", "testcontainerenv",
            "certificate-firendly-name", com.azure.core.util.Context.NONE);
    }
}
