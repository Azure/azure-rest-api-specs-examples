import com.azure.resourcemanager.databricks.models.VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork;

/** Samples for VNetPeering CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/databricks/resource-manager/Microsoft.Databricks/stable/2023-02-01/examples/WorkspaceVirtualNetworkPeeringCreateOrUpdate.json
     */
    /**
     * Sample code: Create vNet Peering for Workspace.
     *
     * @param manager Entry point to AzureDatabricksManager.
     */
    public static void createVNetPeeringForWorkspace(
        com.azure.resourcemanager.databricks.AzureDatabricksManager manager) {
        manager
            .vNetPeerings()
            .define("vNetPeeringTest")
            .withExistingWorkspace("rg", "myWorkspace")
            .withRemoteVirtualNetwork(
                new VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork()
                    .withId(
                        "/subscriptions/0140911e-1040-48da-8bc9-b99fb3dd88a6/resourceGroups/subramantest/providers/Microsoft.Network/virtualNetworks/subramanvnet"))
            .withAllowVirtualNetworkAccess(true)
            .withAllowForwardedTraffic(false)
            .withAllowGatewayTransit(false)
            .withUseRemoteGateways(false)
            .create();
    }
}
