
/**
 * Samples for ManagementLocks DeleteAtResourceLevel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/locks/stable/2020-05-01/examples/
     * ManagementLocks_DeleteAtResourceLevel.json
     */
    /**
     * Sample code: Delete management lock at resource level.
     *
     * @param manager Entry point to ResourceManager.
     */
    public static void
        deleteManagementLockAtResourceLevel(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.managementLockClient().getManagementLocks().deleteAtResourceLevelWithResponse("resourcegroupname",
            "Microsoft.Storage", "parentResourcePath", "storageAccounts", "teststorageaccount", "testlock",
            com.azure.core.util.Context.NONE);
    }
}
