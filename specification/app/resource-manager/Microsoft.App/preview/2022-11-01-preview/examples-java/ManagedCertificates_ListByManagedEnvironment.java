/** Samples for ManagedCertificates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ManagedCertificates_ListByManagedEnvironment.json
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
