
import com.azure.resourcemanager.cosmos.models.NotebookWorkspaceName;

/**
 * Samples for NotebookWorkspaces Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBNotebookWorkspaceStart.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceStart.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBNotebookWorkspaceStart(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getNotebookWorkspaces().start("rg1", "ddb1", NotebookWorkspaceName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
