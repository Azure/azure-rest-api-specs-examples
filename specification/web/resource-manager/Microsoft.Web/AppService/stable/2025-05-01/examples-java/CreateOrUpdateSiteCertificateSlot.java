
import com.azure.resourcemanager.appservice.fluent.models.CertificateInner;
import java.util.Arrays;

/**
 * Samples for SiteCertificates CreateOrUpdateSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/CreateOrUpdateSiteCertificateSlot.json
     */
    /**
     * Sample code: Create Or Update Certificate for slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        createOrUpdateCertificateForSlot(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getSiteCertificates().createOrUpdateSlotWithResponse(
            "testrg123", "testSiteName", "staging", "testc6282", new CertificateInner().withLocation("East US")
                .withPassword("fakeTokenPlaceholder").withHostNames(Arrays.asList("ServerCert")),
            com.azure.core.util.Context.NONE);
    }
}
