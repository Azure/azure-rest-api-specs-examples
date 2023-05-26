/** Samples for ApplicationSecurityGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-11-01/examples/ApplicationSecurityGroupList.json
     */
    /**
     * Sample code: List load balancers in resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listLoadBalancersInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getApplicationSecurityGroups()
            .listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
