
import com.azure.resourcemanager.network.fluent.models.NetworkInterfaceIpConfigurationInner;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkTapInner;

/**
 * Samples for VirtualNetworkTaps CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/VirtualNetworkTapCreate.json
     */
    /**
     * Sample code: Create Virtual Network Tap.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createVirtualNetworkTap(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getVirtualNetworkTaps().createOrUpdate("rg1", "test-vtap",
            new VirtualNetworkTapInner().withLocation("centraluseuap")
                .withDestinationNetworkInterfaceIpConfiguration(new NetworkInterfaceIpConfigurationInner().withId(
                    "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkInterfaces/testNetworkInterface/ipConfigurations/ipconfig1")),
            com.azure.core.util.Context.NONE);
    }
}
