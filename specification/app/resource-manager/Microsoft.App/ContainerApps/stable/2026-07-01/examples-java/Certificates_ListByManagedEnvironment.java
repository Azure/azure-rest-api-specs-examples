
/**
 * Samples for Certificates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Certificates_ListByManagedEnvironment.json
     */
    /**
     * Sample code: List Certificates by Managed Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        listCertificatesByManagedEnvironment(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.certificates().list("examplerg", "testcontainerenv", com.azure.core.util.Context.NONE);
    }
}
