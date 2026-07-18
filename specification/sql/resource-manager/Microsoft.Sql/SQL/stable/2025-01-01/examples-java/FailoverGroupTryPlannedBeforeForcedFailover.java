
/**
 * Samples for FailoverGroups TryPlannedBeforeForcedFailover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/FailoverGroupTryPlannedBeforeForcedFailover.json
     */
    /**
     * Sample code: Try planned before forced failover of a failover group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        tryPlannedBeforeForcedFailoverOfAFailoverGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getFailoverGroups().tryPlannedBeforeForcedFailover("Default",
            "failovergroupsecondaryserver", "failovergrouptest3", com.azure.core.util.Context.NONE);
    }
}
