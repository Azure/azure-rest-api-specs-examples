
import com.azure.core.util.Context;

/** Samples for FailoverGroups ListByServer. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/FailoverGroupList.json
     */
    /**
     * Sample code: List failover group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listFailoverGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getFailoverGroups().listByServer("Default",
            "failover-group-primary-server", Context.NONE);
    }
}
