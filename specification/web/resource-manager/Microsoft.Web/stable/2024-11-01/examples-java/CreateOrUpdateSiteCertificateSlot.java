
import com.azure.resourcemanager.appservice.fluent.models.CertificateInner;
import java.util.Arrays;

/**
 * Samples for SiteCertificates CreateOrUpdateSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/CreateOrUpdateSiteCertificateSlot.
     * json
     */
    /**
     * Sample code: Create Or Update Certificate for slot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateCertificateForSlot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getSiteCertificates().createOrUpdateSlotWithResponse("testrg123",
            "testSiteName", "staging", "testc6282", new CertificateInner().withLocation("East US")
                .withPassword("fakeTokenPlaceholder").withHostNames(Arrays.asList("ServerCert")),
            com.azure.core.util.Context.NONE);
    }
}
