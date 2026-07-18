
/**
 * Samples for InstanceFailoverGroups ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/InstanceFailoverGroupList.json
     */
    /**
     * Sample code: List failover group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listFailoverGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getInstanceFailoverGroups().listByLocation("Default", "Japan East",
            com.azure.core.util.Context.NONE);
    }
}
