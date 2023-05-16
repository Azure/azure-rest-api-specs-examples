/** Samples for ConnectedEnvironmentsCertificates List. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ConnectedEnvironmentsCertificates_ListByConnectedEnvironment.json
     */
    /**
     * Sample code: List Certificates by Connected Environment.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listCertificatesByConnectedEnvironment(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .connectedEnvironmentsCertificates()
            .list("examplerg", "testcontainerenv", com.azure.core.util.Context.NONE);
    }
}
