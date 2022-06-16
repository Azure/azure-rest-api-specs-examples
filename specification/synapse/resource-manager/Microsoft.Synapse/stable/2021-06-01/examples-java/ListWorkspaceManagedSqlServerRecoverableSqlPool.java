import com.azure.core.util.Context;

/** Samples for WorkspaceManagedSqlServerRecoverableSqlPools List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspaceManagedSqlServerRecoverableSqlPool.json
     */
    /**
     * Sample code: List recoverable sql pools for the server.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listRecoverableSqlPoolsForTheServer(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaceManagedSqlServerRecoverableSqlPools().list("wsg-7398", "testWorkspace", Context.NONE);
    }
}
