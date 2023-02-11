/** Samples for SqlPools ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolsInWorkspace.json
     */
    /**
     * Sample code: List SQL Analytics pools in a workspace.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listSQLAnalyticsPoolsInAWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPools().listByWorkspace("sqlcrudtest-6845", "sqlcrudtest-7177", com.azure.core.util.Context.NONE);
    }
}
