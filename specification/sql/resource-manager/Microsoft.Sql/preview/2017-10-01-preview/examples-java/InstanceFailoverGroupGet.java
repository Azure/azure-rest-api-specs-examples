import com.azure.core.util.Context;

/** Samples for InstanceFailoverGroups Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-10-01-preview/examples/InstanceFailoverGroupGet.json
     */
    /**
     * Sample code: Get failover group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFailoverGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getInstanceFailoverGroups()
            .getWithResponse("Default", "Japan East", "failover-group-test", Context.NONE);
    }
}
