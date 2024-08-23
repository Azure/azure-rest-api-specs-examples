
/**
 * Samples for ManagedInstances ListOutboundNetworkDependenciesByManagedInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ListOutboundNetworkDependenciesByManagedInstance.json
     */
    /**
     * Sample code: Gets the collection of outbound network dependencies for the given managed instance.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheCollectionOfOutboundNetworkDependenciesForTheGivenManagedInstance(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances()
            .listOutboundNetworkDependenciesByManagedInstance("sqlcrudtest-7398", "testinstance",
                com.azure.core.util.Context.NONE);
    }
}
