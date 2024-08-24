
/**
 * Samples for InstancePools ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ListInstancePoolsByResourceGroup.json
     */
    /**
     * Sample code: List instance pools by resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listInstancePoolsByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getInstancePools().listByResourceGroup("group1",
            com.azure.core.util.Context.NONE);
    }
}
