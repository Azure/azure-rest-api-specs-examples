
/**
 * Samples for ManagementLocks GetByScope.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/locks/stable/2020-05-01/examples/
     * ManagementLocks_GetAtScope.json
     */
    /**
     * Sample code: Get management lock at scope.
     *
     * @param manager Entry point to ResourceManager.
     */
    public static void getManagementLockAtScope(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.managementLockClient().getManagementLocks().getByScopeWithResponse("subscriptions/subscriptionId",
            "testlock", com.azure.core.util.Context.NONE);
    }
}
