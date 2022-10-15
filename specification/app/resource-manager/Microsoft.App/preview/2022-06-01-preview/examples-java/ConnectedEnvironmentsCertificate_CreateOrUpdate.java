import com.azure.core.util.Context;
import com.azure.resourcemanager.appcontainers.fluent.models.CertificateInner;
import com.azure.resourcemanager.appcontainers.models.CertificateProperties;

/** Samples for ConnectedEnvironmentsCertificates CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ConnectedEnvironmentsCertificate_CreateOrUpdate.json
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
            .createOrUpdateWithResponse(
                "examplerg",
                "testcontainerenv",
                "certificate-firendly-name",
                new CertificateInner()
                    .withLocation("East US")
                    .withProperties(
                        new CertificateProperties()
                            .withPassword("private key password")
                            .withValue("PFX-or-PEM-blob".getBytes())),
                Context.NONE);
    }
}
