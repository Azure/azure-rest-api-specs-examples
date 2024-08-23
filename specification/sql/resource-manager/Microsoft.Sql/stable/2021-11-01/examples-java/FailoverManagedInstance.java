
import com.azure.resourcemanager.sql.models.ReplicaType;

/**
 * Samples for ManagedInstances Failover.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/FailoverManagedInstance.json
     */
    /**
     * Sample code: Failover a managed instance.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void failoverAManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().failover("group1", "instanceName",
            ReplicaType.PRIMARY, com.azure.core.util.Context.NONE);
    }
}
