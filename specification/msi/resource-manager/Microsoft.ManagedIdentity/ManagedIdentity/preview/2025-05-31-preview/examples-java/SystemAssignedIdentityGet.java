
/**
 * Samples for SystemAssignedIdentities GetByScope.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/SystemAssignedIdentityGet.json
     */
    /**
     * Sample code: SystemAssignedIdentityGet.
     * 
     * @param manager Entry point to MsiManager.
     */
    public static void systemAssignedIdentityGet(com.azure.resourcemanager.msi.MsiManager manager) {
        manager.serviceClient().getSystemAssignedIdentities().getByScopeWithResponse(
            "subscriptions/subId/resourceGroups/resourceGroupName/providers/Resource.Provider/resourceType/resourceName/identities/default",
            com.azure.core.util.Context.NONE);
    }
}
