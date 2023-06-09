/** Samples for Workspaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspacesListBySubscription.json
     */
    /**
     * Sample code: Lists workspaces.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void listsWorkspaces(com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager.workspaces().list(com.azure.core.util.Context.NONE);
    }
}
