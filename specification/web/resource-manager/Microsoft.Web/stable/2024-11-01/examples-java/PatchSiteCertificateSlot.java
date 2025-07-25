
import com.azure.resourcemanager.appservice.models.CertificatePatchResource;

/**
 * Samples for SiteCertificates UpdateSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/PatchSiteCertificateSlot.json
     */
    /**
     * Sample code: Patch Certificate for slot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchCertificateForSlot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getSiteCertificates().updateSlotWithResponse("testrg123",
            "testSiteName", "staging", "testc6282",
            new CertificatePatchResource().withKeyVaultId("fakeTokenPlaceholder"), com.azure.core.util.Context.NONE);
    }
}
