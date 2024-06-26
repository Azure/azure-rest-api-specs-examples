
/**
 * Samples for Providers GetAtTenantScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/GetNamedProviderAtTenant.
     * json
     */
    /**
     * Sample code: Get a resource provider at tenant scope.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAResourceProviderAtTenantScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getProviders().getAtTenantScopeWithResponse(
            "Microsoft.Storage", "resourceTypes/aliases", com.azure.core.util.Context.NONE);
    }
}
