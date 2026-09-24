
/**
 * Samples for ManagedInstances ListOutboundNetworkDependenciesByManagedInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ListOutboundNetworkDependenciesByManagedInstance.json
     */
    /**
     * Sample code: Gets the collection of outbound network dependencies for the given managed instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsTheCollectionOfOutboundNetworkDependenciesForTheGivenManagedInstance(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().listOutboundNetworkDependenciesByManagedInstance(
            "sqlcrudtest-7398", "testinstance", com.azure.core.util.Context.NONE);
    }
}
