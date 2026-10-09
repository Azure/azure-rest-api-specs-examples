
/**
 * Samples for AuthorizationOperations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/locks/stable/2020-05-01/examples/
     * ListProviderOperations.json
     */
    /**
     * Sample code: List provider operations.
     *
     * @param manager Entry point to ResourceManager.
     */
    public static void listProviderOperations(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.managementLockClient().getAuthorizationOperations().list(com.azure.core.util.Context.NONE);
    }
}
