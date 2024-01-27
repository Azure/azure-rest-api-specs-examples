
/** Samples for AuthorizationOperations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2017-04-01/examples/
     * ListProviderOperations.json
     */
    /**
     * Sample code: List provider operations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listProviderOperations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().managementLockClient().getAuthorizationOperations()
            .list(com.azure.core.util.Context.NONE);
    }
}
