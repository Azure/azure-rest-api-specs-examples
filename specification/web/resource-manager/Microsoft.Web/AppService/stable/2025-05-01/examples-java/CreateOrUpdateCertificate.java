
import com.azure.resourcemanager.appservice.fluent.models.CertificateInner;
import java.util.Arrays;

/**
 * Samples for Certificates CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/CreateOrUpdateCertificate.json
     */
    /**
     * Sample code: Create Or Update Certificate.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void createOrUpdateCertificate(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getCertificates().createOrUpdateWithResponse(
            "testrg123", "testc6282", new CertificateInner().withLocation("East US")
                .withPassword("fakeTokenPlaceholder").withHostNames(Arrays.asList("ServerCert")),
            com.azure.core.util.Context.NONE);
    }
}
