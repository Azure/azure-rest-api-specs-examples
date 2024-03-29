import com.azure.core.util.Context;

/** Samples for Certificates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/Certificates_ListByManagedEnvironment.json
     */
    /**
     * Sample code: List Certificates by Managed Environment.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listCertificatesByManagedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.certificates().list("examplerg", "testcontainerenv", Context.NONE);
    }
}
