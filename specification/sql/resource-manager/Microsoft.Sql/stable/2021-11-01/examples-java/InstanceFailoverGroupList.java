
import com.azure.core.util.Context;

/** Samples for InstanceFailoverGroups ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/InstanceFailoverGroupList.json
     */
    /**
     * Sample code: List failover group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listFailoverGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getInstanceFailoverGroups().listByLocation("Default", "Japan East",
            Context.NONE);
    }
}
