
/**
 * Samples for SystemAssignedIdentities GetByScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/msi/resource-manager/Microsoft.ManagedIdentity/stable/2024-11-30/examples/SystemAssignedIdentityGet
     * .json
     */
    /**
     * Sample code: MsiOperationsList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void msiOperationsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.identities().manager().serviceClient().getSystemAssignedIdentities().getByScopeWithResponse(
            "subscriptions/subId/resourceGroups/resourceGroupName/providers/Resource.Provider/resourceType/resourceName/identities/default",
            com.azure.core.util.Context.NONE);
    }
}
