import com.azure.core.util.Context;

/** Samples for NetworkSecurityGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/NetworkSecurityGroupDelete.json
     */
    /**
     * Sample code: Delete network security group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteNetworkSecurityGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityGroups().delete("rg1", "testnsg", Context.NONE);
    }
}
