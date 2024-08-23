
/**
 * Samples for InstanceFailoverGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/InstanceFailoverGroupDelete.json
     */
    /**
     * Sample code: Delete failover group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteFailoverGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getInstanceFailoverGroups().delete("Default", "Japan East",
            "failover-group-test-1", com.azure.core.util.Context.NONE);
    }
}
