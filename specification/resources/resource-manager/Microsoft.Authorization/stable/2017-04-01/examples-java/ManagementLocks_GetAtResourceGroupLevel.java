
/** Samples for ManagementLocks GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2017-04-01/examples/
     * ManagementLocks_GetAtResourceGroupLevel.json
     */
    /**
     * Sample code: Get management lock at resource group level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagementLockAtResourceGroupLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().managementLockClient().getManagementLocks()
            .getByResourceGroupWithResponse("resourcegroupname", "testlock", com.azure.core.util.Context.NONE);
    }
}
