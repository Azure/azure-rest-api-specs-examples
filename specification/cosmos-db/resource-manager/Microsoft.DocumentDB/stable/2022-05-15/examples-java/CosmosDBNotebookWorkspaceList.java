import com.azure.core.util.Context;

/** Samples for NotebookWorkspaces ListByDatabaseAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBNotebookWorkspaceList.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceList.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBNotebookWorkspaceList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getNotebookWorkspaces()
            .listByDatabaseAccount("rg1", "ddb1", Context.NONE);
    }
}
