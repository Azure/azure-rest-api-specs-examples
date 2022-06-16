import com.azure.core.util.Context;

/** Samples for SqlPools ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolsInWorkspaceWithFilter.json
     */
    /**
     * Sample code: List SQL Analytics pools in a workspace with filter.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listSQLAnalyticsPoolsInAWorkspaceWithFilter(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPools().listByWorkspace("sqlcrudtest-6845", "sqlcrudtest-7177", Context.NONE);
    }
}
