import com.azure.core.util.Context;

/** Samples for UserAssignedIdentities ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2018-11-30/examples/IdentityListByResourceGroup.json
     */
    /**
     * Sample code: IdentityListByResourceGroup.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void identityListByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .identities()
            .manager()
            .serviceClient()
            .getUserAssignedIdentities()
            .listByResourceGroup("rgName", Context.NONE);
    }
}
