
import com.azure.resourcemanager.network.fluent.models.NetworkSecurityGroupInner;

/**
 * Samples for NetworkSecurityGroups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/NetworkSecurityGroupCreate.
     * json
     */
    /**
     * Sample code: Create network security group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createNetworkSecurityGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getNetworkSecurityGroups().createOrUpdate("rg1", "testnsg",
            new NetworkSecurityGroupInner().withLocation("eastus"), com.azure.core.util.Context.NONE);
    }
}
