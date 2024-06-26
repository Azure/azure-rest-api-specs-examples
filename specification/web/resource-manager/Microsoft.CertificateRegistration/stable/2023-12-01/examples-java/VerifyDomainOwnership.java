
/**
 * Samples for AppServiceCertificateOrders VerifyDomainOwnership.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.CertificateRegistration/stable/2023-12-01/examples/
     * VerifyDomainOwnership.json
     */
    /**
     * Sample code: Verify Domain Ownership.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void verifyDomainOwnership(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceCertificateOrders().verifyDomainOwnershipWithResponse(
            "testrg123", "SampleCertificateOrderName", com.azure.core.util.Context.NONE);
    }
}
