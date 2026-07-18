
/**
 * Samples for FailoverGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FailoverGroupDelete.json
     */
    /**
     * Sample code: Delete failover group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteFailoverGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getFailoverGroups().delete("Default", "failover-group-primary-server",
            "failover-group-test-1", com.azure.core.util.Context.NONE);
    }
}
