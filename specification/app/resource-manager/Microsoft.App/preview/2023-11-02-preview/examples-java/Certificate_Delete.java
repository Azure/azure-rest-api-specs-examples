
/**
 * Samples for Certificates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/Certificate_Delete.json
     */
    /**
     * Sample code: Delete Certificate.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteCertificate(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.certificates().deleteWithResponse("examplerg", "testcontainerenv", "certificate-firendly-name",
            com.azure.core.util.Context.NONE);
    }
}
