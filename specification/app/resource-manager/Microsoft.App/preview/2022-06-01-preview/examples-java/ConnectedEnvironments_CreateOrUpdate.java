import com.azure.resourcemanager.appcontainers.models.CustomDomainConfiguration;

/** Samples for ConnectedEnvironments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ConnectedEnvironments_CreateOrUpdate.json
     */
    /**
     * Sample code: Create kube environments.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createKubeEnvironments(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .connectedEnvironments()
            .define("testenv")
            .withRegion("East US")
            .withExistingResourceGroup("examplerg")
            .withStaticIp("1.2.3.4")
            .withDaprAIConnectionString(
                "InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/")
            .withCustomDomainConfiguration(
                new CustomDomainConfiguration()
                    .withDnsSuffix("www.my-name.com")
                    .withCertificateValue("PFX-or-PEM-blob".getBytes())
                    .withCertificatePassword("private key password".getBytes()))
            .create();
    }
}
