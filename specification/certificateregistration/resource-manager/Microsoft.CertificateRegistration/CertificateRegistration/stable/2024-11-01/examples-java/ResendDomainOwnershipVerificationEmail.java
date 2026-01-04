
import com.azure.resourcemanager.appservice.fluent.models.NameIdentifierInner;

/**
 * Samples for AppServiceCertificateOrders ResendRequestEmails.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/certificateregistration/resource-manager/Microsoft.CertificateRegistration/CertificateRegistration/
     * stable/2024-11-01/examples/ResendDomainOwnershipVerificationEmail.json
     */
    /**
     * Sample code: Resend Domain Ownership verification email.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void resendDomainOwnershipVerificationEmail(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders().resendRequestEmailsWithResponse(
            "testrg123", "SampleCertificateOrderName", new NameIdentifierInner().withName("Domain name"),
            com.azure.core.util.Context.NONE);
    }
}
