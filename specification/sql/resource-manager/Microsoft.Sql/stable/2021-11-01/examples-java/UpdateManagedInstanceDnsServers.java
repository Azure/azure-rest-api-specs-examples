
/**
 * Samples for VirtualClusters UpdateDnsServers.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/UpdateManagedInstanceDnsServers.json
     */
    /**
     * Sample code: Synchronizes the DNS server settings used by the managed instances inside the given virtual cluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void synchronizesTheDNSServerSettingsUsedByTheManagedInstancesInsideTheGivenVirtualCluster(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getVirtualClusters()
            .updateDnsServersWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645", com.azure.core.util.Context.NONE);
    }
}
