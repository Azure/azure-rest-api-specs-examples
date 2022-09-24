import com.azure.core.util.Context;
import com.azure.resourcemanager.cosmos.models.NotebookWorkspaceName;

/** Samples for NotebookWorkspaces Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-08-15/examples/CosmosDBNotebookWorkspaceStart.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceStart.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBNotebookWorkspaceStart(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getNotebookWorkspaces()
            .start("rg1", "ddb1", NotebookWorkspaceName.DEFAULT, Context.NONE);
    }
}
