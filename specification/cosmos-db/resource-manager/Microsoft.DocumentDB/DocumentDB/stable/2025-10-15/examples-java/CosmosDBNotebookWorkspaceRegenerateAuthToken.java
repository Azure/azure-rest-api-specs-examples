
import com.azure.resourcemanager.cosmos.models.NotebookWorkspaceName;

/**
 * Samples for NotebookWorkspaces RegenerateAuthToken.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBNotebookWorkspaceRegenerateAuthToken.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceRegenerateAuthToken.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        cosmosDBNotebookWorkspaceRegenerateAuthToken(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getNotebookWorkspaces().regenerateAuthToken("rg1", "ddb1",
            NotebookWorkspaceName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
