
/**
 * Samples for DistributedAvailabilityGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DistributedAvailabilityGroupsGet.json
     */
    /**
     * Sample code: Gets the distributed availability group info.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsTheDistributedAvailabilityGroupInfo(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDistributedAvailabilityGroups().getWithResponse("testrg",
            "testcl", "dag", com.azure.core.util.Context.NONE);
    }
}
