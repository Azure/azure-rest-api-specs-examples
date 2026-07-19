
/**
 * Samples for DistributedAvailabilityGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DistributedAvailabilityGroupsGet.json
     */
    /**
     * Sample code: Gets the distributed availability group info.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsTheDistributedAvailabilityGroupInfo(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDistributedAvailabilityGroups().getWithResponse("testrg", "testcl", "dag",
            com.azure.core.util.Context.NONE);
    }
}
