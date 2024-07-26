
import com.azure.resourcemanager.appcontainers.fluent.models.CertificateInner;
import com.azure.resourcemanager.appcontainers.models.CertificateProperties;
import com.azure.resourcemanager.appcontainers.models.CertificateType;

/**
 * Samples for Certificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/Certificate_CreateOrUpdate.
     * json
     */
    /**
     * Sample code: Create or Update Certificate.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createOrUpdateCertificate(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.certificates().createOrUpdateWithResponse("examplerg", "testcontainerenv", "certificate-firendly-name",
            new CertificateInner().withLocation("East US")
                .withProperties(new CertificateProperties().withPassword("fakeTokenPlaceholder")
                    .withValue("Y2VydA==".getBytes()).withCertificateType(CertificateType.IMAGE_PULL_TRUSTED_CA)),
            com.azure.core.util.Context.NONE);
    }
}
