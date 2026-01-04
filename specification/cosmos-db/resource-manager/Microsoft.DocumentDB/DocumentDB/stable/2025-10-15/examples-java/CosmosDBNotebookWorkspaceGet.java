
import com.azure.resourcemanager.cosmos.models.NotebookWorkspaceName;

/**
 * Samples for NotebookWorkspaces Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cosmos-db/resource-manager/Microsoft.DocumentDB/DocumentDB/stable/2025-10-15/examples/
     * CosmosDBNotebookWorkspaceGet.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceGet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBNotebookWorkspaceGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cosmosDBAccounts().manager().serviceClient().getNotebookWorkspaces().getWithResponse("rg1", "ddb1",
            NotebookWorkspaceName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
