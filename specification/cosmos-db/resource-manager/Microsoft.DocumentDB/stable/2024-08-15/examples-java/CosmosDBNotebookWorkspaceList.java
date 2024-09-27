
/**
 * Samples for NotebookWorkspaces ListByDatabaseAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2024-08-15/examples/
     * CosmosDBNotebookWorkspaceList.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceList.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBNotebookWorkspaceList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getNotebookWorkspaces().listByDatabaseAccount("rg1", "ddb1",
            com.azure.core.util.Context.NONE);
    }
}
