/** Samples for Workspaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspaceDelete.json
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
