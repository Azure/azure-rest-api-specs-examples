
/** Samples for SystemAssignedIdentities GetByScope. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2023-01-31/examples/SystemAssignedIdentityGet
     * .json
     */
    /**
     * Sample code: MsiOperationsList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void msiOperationsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.identities().manager().serviceClient().getSystemAssignedIdentities().getByScopeWithResponse("scope",
            com.azure.core.util.Context.NONE);
    }
}
