/** Samples for Workspaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/WorkspaceDelete.json
     */
    /**
     * Sample code: Delete a workspace.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void deleteAWorkspace(com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager.workspaces().delete("rg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
