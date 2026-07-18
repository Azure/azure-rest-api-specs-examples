
/**
 * Samples for DistributedAvailabilityGroups ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DistributedAvailabilityGroupsListByInstance.json
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
