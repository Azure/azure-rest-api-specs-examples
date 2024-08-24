
/**
 * Samples for UserAssignedIdentities Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/IdentityDelete.json
     */
    /**
     * Sample code: IdentityDelete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void identityDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.identities().manager().serviceClient().getUserAssignedIdentities().deleteWithResponse("rgName",
            "resourceName", com.azure.core.util.Context.NONE);
    }
}
