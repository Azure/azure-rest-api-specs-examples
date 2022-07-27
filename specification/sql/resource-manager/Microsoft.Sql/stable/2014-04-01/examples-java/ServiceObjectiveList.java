import com.azure.core.util.Context;

/** Samples for ServiceObjectives ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/ServiceObjectiveList.json
     */
    /**
     * Sample code: List service objectives.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listServiceObjectives(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServiceObjectives()
            .listByServer("group1", "sqlcrudtest", Context.NONE);
    }
}
