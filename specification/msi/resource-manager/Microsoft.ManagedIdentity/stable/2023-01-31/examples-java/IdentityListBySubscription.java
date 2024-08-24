
/**
 * Samples for UserAssignedIdentities List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/
     * IdentityListBySubscription.json
     */
    /**
     * Sample code: IdentityListBySubscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void identityListBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.identities().manager().serviceClient().getUserAssignedIdentities().list(com.azure.core.util.Context.NONE);
    }
}
