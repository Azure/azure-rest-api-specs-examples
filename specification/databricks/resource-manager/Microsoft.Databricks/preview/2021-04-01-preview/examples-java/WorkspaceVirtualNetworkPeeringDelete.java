/** Samples for VNetPeering Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/WorkspaceVirtualNetworkPeeringDelete.json
     */
    /**
     * Sample code: Delete a workspace vNet Peering.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void deleteAWorkspaceVNetPeering(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager.vNetPeerings().delete("rg", "myWorkspace", "vNetPeering", com.azure.core.util.Context.NONE);
    }
}
