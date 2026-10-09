
/**
 * Samples for ManagementLocks List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/locks/stable/2020-05-01/examples/
     * ManagementLocks_ListAtSubscriptionLevel.json
     */
    /**
     * Sample code: List management locks at subscription level.
     *
     * @param manager Entry point to ResourceManager.
     */
    public static void
        listManagementLocksAtSubscriptionLevel(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.managementLockClient().getManagementLocks().list(null, com.azure.core.util.Context.NONE);
    }
}
