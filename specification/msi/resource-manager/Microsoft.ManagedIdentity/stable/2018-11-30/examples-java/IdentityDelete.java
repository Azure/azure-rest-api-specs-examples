import com.azure.core.util.Context;

/** Samples for UserAssignedIdentities Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2018-11-30/examples/IdentityDelete.json
     */
    /**
     * Sample code: IdentityDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void identityDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .identities()
            .manager()
            .serviceClient()
            .getUserAssignedIdentities()
            .deleteWithResponse("rgName", "resourceName", Context.NONE);
    }
}
