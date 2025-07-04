
import com.azure.resourcemanager.appservice.fluent.models.DomainOwnershipIdentifierInner;

/**
 * Samples for Domains UpdateOwnershipIdentifier.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2024-11-01/examples/
     * UpdateAppServiceDomainOwnershipIdentifier.json
     */
    /**
     * Sample code: Update App Service Domain OwnershipIdentifier.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAppServiceDomainOwnershipIdentifier(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().updateOwnershipIdentifierWithResponse("testrg123",
            "example.com", "SampleOwnershipId",
            new DomainOwnershipIdentifierInner().withOwnershipId("SampleOwnershipId"),
            com.azure.core.util.Context.NONE);
    }
}
