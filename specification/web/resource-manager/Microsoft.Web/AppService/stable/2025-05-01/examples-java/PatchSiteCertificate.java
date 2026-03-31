
import com.azure.resourcemanager.appservice.models.CertificatePatchResource;

/**
 * Samples for SiteCertificates Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/PatchSiteCertificate.json
     */
    /**
     * Sample code: Patch Certificate.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void patchCertificate(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getSiteCertificates().updateWithResponse("testrg123", "testSiteName", "testc6282",
            new CertificatePatchResource().withKeyVaultId("fakeTokenPlaceholder"), com.azure.core.util.Context.NONE);
    }
}
