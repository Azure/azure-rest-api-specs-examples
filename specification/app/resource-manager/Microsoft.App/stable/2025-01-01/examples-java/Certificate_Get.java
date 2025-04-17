
/**
 * Samples for Certificates Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/Certificate_Get.json
     */
    /**
     * Sample code: Get Certificate.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getCertificate(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.certificates().getWithResponse("examplerg", "testcontainerenv", "certificate-firendly-name",
            com.azure.core.util.Context.NONE);
    }
}
