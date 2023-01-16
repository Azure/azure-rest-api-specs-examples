/** Samples for VNetPeering ListByWorkspace. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/preview/2021-04-01-preview/examples/WorkspaceVirtualNetPeeringList.json
     */
    /**
     * Sample code: List all vNet Peerings for the workspace.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void listAllVNetPeeringsForTheWorkspace(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager.vNetPeerings().listByWorkspace("rg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
