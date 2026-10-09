
/**
 * Samples for ManagementLocks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/locks/stable/2020-05-01/examples/
     * ManagementLocks_GetAtResourceGroupLevel.json
     */
    /**
     * Sample code: Get management lock at resource group level.
     *
     * @param manager Entry point to ResourceManager.
     */
    public static void
        getManagementLockAtResourceGroupLevel(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.managementLockClient().getManagementLocks().getByResourceGroupWithResponse("resourcegroupname",
            "testlock", com.azure.core.util.Context.NONE);
    }
}
