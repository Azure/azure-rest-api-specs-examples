
import com.azure.resourcemanager.network.models.Nat64Capability;
import com.azure.resourcemanager.network.models.VirtualNetworkApplianceCapabilityProperties;

/**
 * Samples for VirtualNetworkApplianceCapabilities CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/VirtualNetworkApplianceCapabilities_CreateOrUpdate_NAT64.json
     */
    /**
     * Sample code: Create virtual network appliance NAT64 capability.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void
        createVirtualNetworkApplianceNAT64Capability(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkApplianceCapabilities().createOrUpdate("rg1", "test-vna", "nat64",
            new Nat64Capability().withProperties(new VirtualNetworkApplianceCapabilityProperties()),
            com.azure.core.util.Context.NONE);
    }
}
