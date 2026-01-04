
import com.azure.resourcemanager.cosmos.models.NotebookWorkspaceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.NotebookWorkspaceName;

/**
 * Samples for NotebookWorkspaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBNotebookWorkspaceCreate.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBNotebookWorkspaceCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getNotebookWorkspaces().createOrUpdate("rg1", "ddb1",
            NotebookWorkspaceName.DEFAULT, new NotebookWorkspaceCreateUpdateParameters(),
            com.azure.core.util.Context.NONE);
    }
}
