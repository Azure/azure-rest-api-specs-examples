
/** Samples for UserAssignedIdentities GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/IdentityGet.json
     */
    /**
     * Sample code: IdentityGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void identityGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.identities().manager().serviceClient().getUserAssignedIdentities()
            .getByResourceGroupWithResponse("rgName", "resourceName", com.azure.core.util.Context.NONE);
    }
}
