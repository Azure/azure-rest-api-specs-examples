
/**
 * Samples for DistributedAvailabilityGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DistributedAvailabilityGroupsDelete.
     * json
     */
    /**
     * Sample code: Initiate a distributed availability group drop.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void initiateADistributedAvailabilityGroupDrop(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDistributedAvailabilityGroups().delete("testrg", "testcl",
            "dag", com.azure.core.util.Context.NONE);
    }
}
