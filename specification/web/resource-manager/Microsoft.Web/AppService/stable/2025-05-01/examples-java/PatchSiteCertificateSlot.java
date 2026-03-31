
import com.azure.resourcemanager.appservice.models.CertificatePatchResource;

/**
 * Samples for SiteCertificates UpdateSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/PatchSiteCertificateSlot.json
     */
    /**
     * Sample code: Patch Certificate for slot.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void patchCertificateForSlot(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getSiteCertificates().updateSlotWithResponse("testrg123", "testSiteName", "staging",
            "testc6282", new CertificatePatchResource().withKeyVaultId("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
