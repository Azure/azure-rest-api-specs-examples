
/**
 * Samples for InstanceFailoverGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/InstanceFailoverGroupDelete.json
     */
    /**
     * Sample code: Delete failover group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteFailoverGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getInstanceFailoverGroups().delete("Default", "Japan East", "failover-group-test-1",
            com.azure.core.util.Context.NONE);
    }
}
