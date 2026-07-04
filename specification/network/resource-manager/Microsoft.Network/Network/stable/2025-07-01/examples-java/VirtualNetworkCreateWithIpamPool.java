
import com.azure.resourcemanager.network.fluent.models.SubnetInner;
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkInner;
import com.azure.resourcemanager.network.models.AddressSpace;
import com.azure.resourcemanager.network.models.IpamPoolPrefixAllocation;
import java.util.Arrays;

/**
 * Samples for VirtualNetworks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkCreateWithIpamPool.json
     */
    /**
     * Sample code: Create virtual network with ipamPool.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createVirtualNetworkWithIpamPool(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworks().createOrUpdate("rg1", "test-vnet", new VirtualNetworkInner()
            .withLocation("eastus")
            .withAddressSpace(new AddressSpace().withIpamPoolPrefixAllocations(
                Arrays.asList(new IpamPoolPrefixAllocation().withNumberOfIpAddresses("65536").withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/nm1/ipamPools/testIpamPool"))))
            .withSubnets(Arrays.asList(new SubnetInner().withName("test-1").withIpamPoolPrefixAllocations(
                Arrays.asList(new IpamPoolPrefixAllocation().withNumberOfIpAddresses("80").withId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Network/networkManagers/nm1/ipamPools/testIpamPool"))))),
            com.azure.core.util.Context.NONE);
    }
}
