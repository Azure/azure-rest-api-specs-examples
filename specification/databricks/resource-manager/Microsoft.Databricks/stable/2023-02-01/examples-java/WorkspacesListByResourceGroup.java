/** Samples for Workspaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspacesListByResourceGroup.json
     */
    /**
     * Sample code: Lists workspaces.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void listsWorkspaces(com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager.workspaces().listByResourceGroup("rg", com.azure.core.util.Context.NONE);
    }
}
