
/**
 * Samples for ManagedCertificates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * ManagedCertificate_Delete.json
     */
    /**
     * Sample code: Delete Certificate.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteCertificate(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedCertificates().deleteWithResponse("examplerg", "testcontainerenv", "certificate-firendly-name",
            com.azure.core.util.Context.NONE);
    }
}
