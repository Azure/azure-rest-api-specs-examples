
import com.azure.resourcemanager.network.models.PLIPForwardersCapability;
import com.azure.resourcemanager.network.models.VirtualNetworkApplianceCapabilityIpVersion;
import com.azure.resourcemanager.network.models.VirtualNetworkApplianceCapabilityProperties;

/**
 * Samples for VirtualNetworkApplianceCapabilities CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/VirtualNetworkApplianceCapabilities_CreateOrUpdate_PLIPForwarders.json
     */
    /**
     * Sample code: Create virtual network appliance PLIPForwarders capability.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createVirtualNetworkAppliancePLIPForwardersCapability(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkApplianceCapabilities().createOrUpdate("rg1", "test-vna",
            "pl-ipforwarders",
            new PLIPForwardersCapability().withProperties(new VirtualNetworkApplianceCapabilityProperties()
                .withIpVersion(VirtualNetworkApplianceCapabilityIpVersion.IPV6)),
            com.azure.core.util.Context.NONE);
    }
}
