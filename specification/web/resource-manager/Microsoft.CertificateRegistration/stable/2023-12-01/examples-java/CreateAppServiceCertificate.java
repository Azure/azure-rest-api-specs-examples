
import com.azure.resourcemanager.appservice.fluent.models.AppServiceCertificateResourceInner;

/**
 * Samples for AppServiceCertificateOrders CreateOrUpdateCertificate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/
     * CreateAppServiceCertificate.json
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
