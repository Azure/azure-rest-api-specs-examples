
/**
 * Samples for ManagedCertificates List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ManagedCertificates_ListByManagedEnvironment.json
     */
    /**
     * Sample code: List Managed Certificates by Managed Environment.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listManagedCertificatesByManagedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.managedCertificates().list("examplerg", "testcontainerenv", com.azure.core.util.Context.NONE);
    }
}
