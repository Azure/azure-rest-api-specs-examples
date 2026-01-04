
/**
 * Samples for Domains GetOwnershipIdentifier.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/domainregistration/resource-manager/Microsoft.DomainRegistration/DomainRegistration/stable/2024-11-
     * 01/examples/GetDomainOwnershipIdentifier.json
     */
    /**
     * Sample code: Get Domain Ownership Identifier.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDomainOwnershipIdentifier(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().getOwnershipIdentifierWithResponse("testrg123",
            "example.com", "SampleOwnershipId", com.azure.core.util.Context.NONE);
    }
}
