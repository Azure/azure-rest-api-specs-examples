
/**
 * Samples for ManagementLocks DeleteAtSubscriptionLevel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/locks/stable/2020-05-01/examples/
     * ManagementLocks_DeleteAtSubscriptionLevel.json
     */
    /**
     * Sample code: Delete management lock at subscription level.
     *
     * @param manager Entry point to ResourceManager.
     */
    public static void
        deleteManagementLockAtSubscriptionLevel(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.managementLockClient().getManagementLocks().deleteAtSubscriptionLevelWithResponse("testlock",
            com.azure.core.util.Context.NONE);
    }
}
