
import com.azure.resourcemanager.appservice.fluent.models.AppServiceCertificatePatchResourceInner;

/**
 * Samples for AppServiceCertificateOrders UpdateCertificate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/
     * UpdateAppServiceCertificate.json
     */
    /**
     * Sample code: Update Certificate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders().updateCertificateWithResponse(
            "testrg123", "SampleCertificateOrderName", "SampleCertName1", new AppServiceCertificatePatchResourceInner()
                .withKeyVaultId("fakeTokenPlaceholder").withKeyVaultSecretName("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
