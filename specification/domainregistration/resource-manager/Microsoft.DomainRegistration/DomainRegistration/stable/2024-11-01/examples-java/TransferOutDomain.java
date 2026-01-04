
/**
 * Samples for Domains TransferOut.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/domainregistration/resource-manager/Microsoft.DomainRegistration/DomainRegistration/stable/2024-11-
     * 01/examples/TransferOutDomain.json
     */
    /**
     * Sample code: Transfer out domain.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void transferOutDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().transferOutWithResponse("testrg123", "example.com",
            com.azure.core.util.Context.NONE);
    }
}
