
/**
 * Samples for NetworkFunctionDefinitionVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * AzureOperatorNexus/VirtualNetworkFunctionDefinitionVersionDelete.json
     */
    /**
     * Sample code: Delete a network function definition version for AzureOperatorNexus VNF.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void deleteANetworkFunctionDefinitionVersionForAzureOperatorNexusVNF(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.networkFunctionDefinitionVersions().delete("rg", "TestPublisher",
            "TestNetworkFunctionDefinitionGroupName", "1.0.0", com.azure.core.util.Context.NONE);
    }
}
