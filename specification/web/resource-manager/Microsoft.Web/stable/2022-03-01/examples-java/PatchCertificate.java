import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.models.CertificatePatchResource;

/** Samples for Certificates Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-03-01/examples/PatchCertificate.json
     */
    /**
     * Sample code: Patch Certificate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void patchCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getCertificates()
            .updateWithResponse(
                "testrg123", "testc6282", new CertificatePatchResource().withPassword("<password>"), Context.NONE);
    }
}
