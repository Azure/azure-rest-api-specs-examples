
import com.azure.resourcemanager.appservice.fluent.models.AppServiceCertificateResourceInner;

/**
 * Samples for AppServiceCertificateOrders CreateOrUpdateCertificate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/certificateregistration/resource-manager/Microsoft.CertificateRegistration/CertificateRegistration/
     * stable/2024-11-01/examples/CreateAppServiceCertificate.json
     */
    /**
     * Sample code: Create Certificate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createCertificate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders().createOrUpdateCertificate(
            "testrg123", "SampleCertificateOrderName", "SampleCertName1",
            new AppServiceCertificateResourceInner().withLocation("Global").withKeyVaultId("fakeTokenPlaceholder")
                .withKeyVaultSecretName("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
