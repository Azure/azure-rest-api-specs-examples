
/**
 * Samples for DistributedAvailabilityGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DistributedAvailabilityGroupsDelete.json
     */
    /**
     * Sample code: Initiate a distributed availability group drop.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        initiateADistributedAvailabilityGroupDrop(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDistributedAvailabilityGroups().delete("testrg", "testcl", "dag",
            com.azure.core.util.Context.NONE);
    }
}
