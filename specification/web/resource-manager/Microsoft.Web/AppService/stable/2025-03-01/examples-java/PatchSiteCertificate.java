
import com.azure.resourcemanager.appservice.models.CertificatePatchResource;

/**
 * Samples for SiteCertificates Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/PatchSiteCertificate.json
     */
    /**
     * Sample code: Patch Certificate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getSiteCertificates().updateWithResponse("testrg123", "testSiteName",
            "testc6282", new CertificatePatchResource().withKeyVaultId("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
