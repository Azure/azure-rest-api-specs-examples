
import com.azure.resourcemanager.cosmos.models.NotebookWorkspaceCreateUpdateParameters;
import com.azure.resourcemanager.cosmos.models.NotebookWorkspaceName;

/**
 * Samples for NotebookWorkspaces CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15/CosmosDBNotebookWorkspaceCreate.json
     */
    /**
     * Sample code: CosmosDBNotebookWorkspaceCreate.
     * 
     * @param manager Entry point to CosmosManager.
     */
    public static void cosmosDBNotebookWorkspaceCreate(com.azure.resourcemanager.cosmos.CosmosManager manager) {
        manager.serviceClient().getNotebookWorkspaces().createOrUpdate("rg1", "ddb1", NotebookWorkspaceName.DEFAULT,
            new NotebookWorkspaceCreateUpdateParameters(), com.azure.core.util.Context.NONE);
    }
}
