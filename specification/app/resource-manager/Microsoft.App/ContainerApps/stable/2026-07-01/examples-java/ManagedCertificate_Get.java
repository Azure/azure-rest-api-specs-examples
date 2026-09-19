
/**
 * Samples for ManagedCertificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedCertificate_Get.json
     */
    /**
     * Sample code: Get Certificate.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getCertificate(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedCertificates().getWithResponse("examplerg", "testcontainerenv", "certificate-firendly-name",
            com.azure.core.util.Context.NONE);
    }
}
