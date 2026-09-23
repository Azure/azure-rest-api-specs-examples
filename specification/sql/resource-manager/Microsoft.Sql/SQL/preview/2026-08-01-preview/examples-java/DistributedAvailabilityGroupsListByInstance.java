
/**
 * Samples for DistributedAvailabilityGroups ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/DistributedAvailabilityGroupsListByInstance.json
     */
    /**
     * Sample code: Lists all distributed availability groups in instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listsAllDistributedAvailabilityGroupsInInstance(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDistributedAvailabilityGroups().listByInstance("testrg", "testcl",
            com.azure.core.util.Context.NONE);
    }
}
