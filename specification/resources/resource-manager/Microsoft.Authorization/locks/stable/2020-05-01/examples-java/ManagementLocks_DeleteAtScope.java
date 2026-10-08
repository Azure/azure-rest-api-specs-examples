
/**
 * Samples for ManagementLocks DeleteByScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/locks/stable/2020-05-01/examples/
     * ManagementLocks_DeleteAtScope.json
     */
    /**
     * Sample code: Delete management lock at scope.
     *
     * @param manager Entry point to ResourceManager.
     */
    public static void deleteManagementLockAtScope(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.managementLockClient().getManagementLocks().deleteByScopeWithResponse("subscriptions/subscriptionId",
            "testlock", com.azure.core.util.Context.NONE);
    }
}
