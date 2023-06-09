/** Samples for Workspaces GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspaceGet.json
     */
    /**
     * Sample code: Get a workspace.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void getAWorkspace(com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("rg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
