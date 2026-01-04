
import com.azure.resourcemanager.appservice.fluent.models.DomainOwnershipIdentifierInner;

/**
 * Samples for Domains CreateOrUpdateOwnershipIdentifier.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/domainregistration/resource-manager/Microsoft.DomainRegistration/DomainRegistration/stable/2024-11-
     * 01/examples/CreateAppServiceDomainOwnershipIdentifier.json
     */
    /**
     * Sample code: Create App Service Domain OwnershipIdentifier.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAppServiceDomainOwnershipIdentifier(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().createOrUpdateOwnershipIdentifierWithResponse(
            "testrg123", "example.com", "SampleOwnershipId",
            new DomainOwnershipIdentifierInner().withOwnershipId("SampleOwnershipId"),
            com.azure.core.util.Context.NONE);
    }
}
