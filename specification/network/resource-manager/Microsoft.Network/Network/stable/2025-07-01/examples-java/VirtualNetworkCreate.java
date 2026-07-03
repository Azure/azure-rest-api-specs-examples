
import com.azure.resourcemanager.network.fluent.models.VirtualNetworkInner;
import com.azure.resourcemanager.network.models.AddressSpace;
import java.util.Arrays;

/**
 * Samples for VirtualNetworks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/VirtualNetworkCreate.json
     */
    /**
     * Sample code: Create virtual network.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createVirtualNetwork(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworks().createOrUpdate("rg1", "test-vnet",
            new VirtualNetworkInner().withLocation("eastus")
                .withAddressSpace(new AddressSpace().withAddressPrefixes(Arrays.asList("10.0.0.0/16")))
                .withFlowTimeoutInMinutes(10),
            com.azure.core.util.Context.NONE);
    }
}
