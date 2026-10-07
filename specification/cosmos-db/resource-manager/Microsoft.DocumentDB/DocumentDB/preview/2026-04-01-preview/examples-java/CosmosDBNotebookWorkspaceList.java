
/**
 * Samples for NotebookWorkspaces ListByDatabaseAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/CosmosDBNotebookWorkspaceList.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceList.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBNotebookWorkspaceList(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getNotebookWorkspaces().listByDatabaseAccount("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
