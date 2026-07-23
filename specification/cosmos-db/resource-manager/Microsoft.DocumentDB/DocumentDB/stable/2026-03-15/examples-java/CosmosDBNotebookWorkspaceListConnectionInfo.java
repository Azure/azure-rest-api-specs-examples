
import com.azure.resourcemanager.cosmos.models.NotebookWorkspaceName;

/**
 * Samples for NotebookWorkspaces ListConnectionInfo.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBNotebookWorkspaceListConnectionInfo.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceListConnectionInfo.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void
        cosmosDBNotebookWorkspaceListConnectionInfo(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getNotebookWorkspaces().listConnectionInfoWithResponse("rg1", "ddb1",
            NotebookWorkspaceName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
