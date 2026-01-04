
/**
 * Samples for Domains DeleteOwnershipIdentifier.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/domainregistration/resource-manager/Microsoft.DomainRegistration/DomainRegistration/stable/2024-11-
     * 01/examples/DeleteAppServiceDomainOwnershipIdentifier.json
     */
    /**
     * Sample code: Delete App Service Domain Ownership Identifier.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAppServiceDomainOwnershipIdentifier(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().deleteOwnershipIdentifierWithResponse("testrg123",
            "example.com", "ownershipIdentifier", com.azure.core.util.Context.NONE);
    }
}
