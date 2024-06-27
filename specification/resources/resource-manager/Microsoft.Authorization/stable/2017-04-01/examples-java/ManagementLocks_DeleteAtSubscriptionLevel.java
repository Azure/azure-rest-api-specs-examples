
/**
 * Samples for ManagementLocks DeleteAtSubscriptionLevel.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2017-04-01/examples/
     * ManagementLocks_DeleteAtSubscriptionLevel.json
     */
    /**
     * Sample code: Delete management lock at subscription level.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteManagementLockAtSubscriptionLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().managementLockClient().getManagementLocks()
            .deleteAtSubscriptionLevelWithResponse("testlock", com.azure.core.util.Context.NONE);
    }
}
