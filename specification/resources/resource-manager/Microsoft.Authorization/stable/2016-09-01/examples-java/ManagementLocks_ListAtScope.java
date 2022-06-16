import com.azure.core.util.Context;

/** Samples for ManagementLocks ListByScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2016-09-01/examples/ManagementLocks_ListAtScope.json
     */
    /**
     * Sample code: List management locks at scope.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagementLocksAtScope(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .managementLockClient()
            .getManagementLocks()
            .listByScope("subscriptions/subscriptionId", null, Context.NONE);
    }
}
