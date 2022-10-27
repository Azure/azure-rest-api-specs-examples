import com.azure.core.util.Context;

/** Samples for LoadBalancers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/LoadBalancerList.json
     */
    /**
     * Sample code: List load balancers in resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listLoadBalancersInResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getLoadBalancers().listByResourceGroup("rg1", Context.NONE);
    }
}
