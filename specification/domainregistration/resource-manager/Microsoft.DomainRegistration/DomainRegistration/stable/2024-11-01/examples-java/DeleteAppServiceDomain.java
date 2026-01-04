
/**
 * Samples for Domains Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/domainregistration/resource-manager/Microsoft.DomainRegistration/DomainRegistration/stable/2024-11-
     * 01/examples/DeleteAppServiceDomain.json
     */
    /**
     * Sample code: Delete App Service Domain.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAppServiceDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().deleteWithResponse("testrg123", "example.com", true,
            com.azure.core.util.Context.NONE);
    }
}
