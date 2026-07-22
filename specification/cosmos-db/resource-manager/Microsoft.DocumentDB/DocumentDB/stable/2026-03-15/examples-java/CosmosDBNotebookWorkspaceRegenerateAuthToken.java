
import com.azure.resourcemanager.cosmos.models.NotebookWorkspaceName;

/**
 * Samples for NotebookWorkspaces RegenerateAuthToken.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBNotebookWorkspaceRegenerateAuthToken.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceRegenerateAuthToken.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBNotebookWorkspaceRegenerateAuthToken(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getNotebookWorkspaces().regenerateAuthToken("rg1", "ddb1",
            NotebookWorkspaceName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
