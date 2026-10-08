
/**
 * Samples for ManagementLocks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/locks/stable/2020-05-01/examples/
     * ManagementLocks_ListAtResourceGroupLevel.json
     */
    /**
     * Sample code: List management groups at resource group level.
     *
     * @param manager Entry point to ResourceManager.
     */
    public static void
        listManagementGroupsAtResourceGroupLevel(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.managementLockClient().getManagementLocks().listByResourceGroup("resourcegroupname", null,
            com.azure.core.util.Context.NONE);
    }
}
