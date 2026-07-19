
/**
 * Samples for SynapseLinkWorkspaces ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SynapseLinkWorkspaceListByDatabase.json
     */
    /**
     * Sample code: List all synapselink workspaces for the given database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listAllSynapselinkWorkspacesForTheGivenDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSynapseLinkWorkspaces().listByDatabase("Default-SQL-SouthEastAsia", "testsvr",
            "dbSynapse", com.azure.core.util.Context.NONE);
    }
}
