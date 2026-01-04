
/**
 * Samples for DomainRegistrationProvider ListOperations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/domainregistration/resource-manager/Microsoft.DomainRegistration/DomainRegistration/stable/2024-11-
     * 01/examples/ListOperations.json
     */
    /**
     * Sample code: List operations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listOperations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomainRegistrationProviders()
            .listOperations(com.azure.core.util.Context.NONE);
    }
}
