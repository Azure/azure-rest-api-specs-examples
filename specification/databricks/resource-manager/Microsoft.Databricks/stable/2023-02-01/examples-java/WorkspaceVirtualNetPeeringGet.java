/** Samples for VNetPeering Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspaceVirtualNetPeeringGet.json
     */
    /**
     * Sample code: Get a workspace with vNet Peering Configured.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void getAWorkspaceWithVNetPeeringConfigured(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager.vNetPeerings().getWithResponse("rg", "myWorkspace", "vNetPeering", com.azure.core.util.Context.NONE);
    }
}
