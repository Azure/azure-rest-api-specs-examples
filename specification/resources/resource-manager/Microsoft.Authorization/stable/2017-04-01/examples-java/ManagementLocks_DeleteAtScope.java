
/** Samples for ManagementLocks DeleteByScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2017-04-01/examples/
     * ManagementLocks_DeleteAtScope.json
     */
    /**
     * Sample code: Delete management lock at scope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteManagementLockAtScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().managementLockClient().getManagementLocks()
            .deleteByScopeWithResponse("subscriptions/subscriptionId", "testlock", com.azure.core.util.Context.NONE);
    }
}
