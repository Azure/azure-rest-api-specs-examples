
/**
 * Samples for WorkspaceManagedSqlServerRecoverableSqlPools Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * GetWorkspaceManagedSqlServerRecoverableSqlPool.json
     */
    /**
     * Sample code: Get recoverable sql pools for the server.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void getRecoverableSqlPoolsForTheServer(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaceManagedSqlServerRecoverableSqlPools().getWithResponse("wsg-7398", "testWorkspace",
            "recoverableSqlpools-1235", com.azure.core.util.Context.NONE);
    }
}
