import com.azure.core.util.Context;

/** Samples for AppServiceCertificateOrders ResendEmail. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2022-03-01/examples/ResendAppServiceCertificateOrderEmail.json
     */
    /**
     * Sample code: Resend App Service Certificate Order email.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resendAppServiceCertificateOrderEmail(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServiceCertificateOrders()
            .resendEmailWithResponse("testrg123", "SampleCertificateOrderName", Context.NONE);
    }
}
