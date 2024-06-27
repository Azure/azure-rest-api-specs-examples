
/**
 * Samples for Domains ListOwnershipIdentifiers.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2023-12-01/examples/
     * ListDomainOwnershipIdentifiers.json
     */
    /**
     * Sample code: List Domain Ownership Identifiers.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDomainOwnershipIdentifiers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().listOwnershipIdentifiers("testrg123", "example.com",
            com.azure.core.util.Context.NONE);
    }
}
