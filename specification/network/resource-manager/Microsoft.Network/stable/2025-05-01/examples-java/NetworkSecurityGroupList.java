
/**
 * Samples for NetworkSecurityGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkSecurityGroupList.json
     */
    /**
     * Sample code: List network security groups in resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listNetworkSecurityGroupsInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityGroups().listByResourceGroup("rg1",
            com.azure.core.util.Context.NONE);
    }
}
