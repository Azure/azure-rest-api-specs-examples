
import com.azure.resourcemanager.appcontainers.fluent.models.CertificateInner;
import com.azure.resourcemanager.appcontainers.models.CertificateProperties;

/**
 * Samples for Certificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/Certificate_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update Certificate.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        createOrUpdateCertificate(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.certificates().createOrUpdateWithResponse("examplerg", "testcontainerenv", "certificate-firendly-name",
            new CertificateInner().withLocation("East US").withProperties(
                new CertificateProperties().withPassword("fakeTokenPlaceholder").withValue("Y2VydA==".getBytes())),
            com.azure.core.util.Context.NONE);
    }
}
