import com.azure.resourcemanager.appcontainers.models.CertificateProperties;

/** Samples for ConnectedEnvironmentsCertificates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ConnectedEnvironmentsCertificate_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update Certificate.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateCertificate(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .connectedEnvironmentsCertificates()
            .define("certificate-firendly-name")
            .withRegion("East US")
            .withExistingConnectedEnvironment("examplerg", "testcontainerenv")
            .withProperties(
                new CertificateProperties().withPassword("fakeTokenPlaceholder").withValue("Y2VydA==".getBytes()))
            .create();
    }
}
