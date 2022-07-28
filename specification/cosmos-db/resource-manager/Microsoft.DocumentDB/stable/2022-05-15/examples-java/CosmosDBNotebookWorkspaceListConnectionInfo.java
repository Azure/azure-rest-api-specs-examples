import com.azure.core.util.Context;
import com.azure.resourcemanager.cosmos.models.NotebookWorkspaceName;

/** Samples for NotebookWorkspaces ListConnectionInfo. */
public final class Main {
    /*
     * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2022-05-15/examples/CosmosDBNotebookWorkspaceListConnectionInfo.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceListConnectionInfo.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cosmosDBNotebookWorkspaceListConnectionInfo(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cosmosDBAccounts()
            .manager()
            .serviceClient()
            .getNotebookWorkspaces()
            .listConnectionInfoWithResponse("rg1", "ddb1", NotebookWorkspaceName.DEFAULT, Context.NONE);
    }
}
