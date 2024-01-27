
/** Samples for ManagementLocks GetByScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2017-04-01/examples/
     * ManagementLocks_GetAtScope.json
     */
    /**
     * Sample code: Get management lock at scope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagementLockAtScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().managementLockClient().getManagementLocks()
            .getByScopeWithResponse("subscriptions/subscriptionId", "testlock", com.azure.core.util.Context.NONE);
    }
}
