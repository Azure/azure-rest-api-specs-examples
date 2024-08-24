
/**
 * Samples for FailoverGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/FailoverGroupGet.json
     */
    /**
     * Sample code: Get failover group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFailoverGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getFailoverGroups().getWithResponse("Default",
            "failover-group-primary-server", "failover-group-test", com.azure.core.util.Context.NONE);
    }
}
