
/** Samples for ManagementLocks DeleteAtResourceLevel. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2017-04-01/examples/
     * ManagementLocks_DeleteAtResourceLevel.json
     */
    /**
     * Sample code: Delete management lock at resource level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteManagementLockAtResourceLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().managementLockClient().getManagementLocks()
            .deleteAtResourceLevelWithResponse("resourcegroupname", "Microsoft.Storage", "parentResourcePath",
                "storageAccounts", "teststorageaccount", "testlock", com.azure.core.util.Context.NONE);
    }
}
