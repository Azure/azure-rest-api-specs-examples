
import com.azure.resourcemanager.sql.models.ReplicaType;

/**
 * Samples for ManagedInstances Failover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/FailoverManagedInstance.json
     */
    /**
     * Sample code: Failover a managed instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void failoverAManagedInstance(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().failover("group1", "instanceName", ReplicaType.PRIMARY,
            com.azure.core.util.Context.NONE);
    }
}
