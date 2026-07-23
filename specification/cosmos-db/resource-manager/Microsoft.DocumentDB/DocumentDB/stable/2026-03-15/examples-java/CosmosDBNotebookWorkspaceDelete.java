
import com.azure.resourcemanager.cosmos.models.NotebookWorkspaceName;

/**
 * Samples for NotebookWorkspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBNotebookWorkspaceDelete.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceDelete.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBNotebookWorkspaceDelete(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getNotebookWorkspaces().delete("rg1", "ddb1", NotebookWorkspaceName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
