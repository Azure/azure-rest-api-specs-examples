
/**
 * Samples for FailoverGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/FailoverGroupDelete.json
     */
    /**
     * Sample code: Delete failover group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteFailoverGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getFailoverGroups().delete("Default",
            "failover-group-primary-server", "failover-group-test-1", com.azure.core.util.Context.NONE);
    }
}
