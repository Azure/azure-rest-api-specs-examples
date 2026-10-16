
import com.azure.resourcemanager.network.models.PLGatewayFastpathCapability;
import com.azure.resourcemanager.network.models.VirtualNetworkApplianceCapabilityIpVersion;
import com.azure.resourcemanager.network.models.VirtualNetworkApplianceCapabilityProperties;

/**
 * Samples for VirtualNetworkApplianceCapabilities CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/VirtualNetworkApplianceCapabilities_CreateOrUpdate.json
     */
    /**
     * Sample code: Create virtual network appliance PLGatewayFastpath capability.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createVirtualNetworkAppliancePLGatewayFastpathCapability(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getVirtualNetworkApplianceCapabilities().createOrUpdate("rg1", "test-vna",
            "pl-fastpath",
            new PLGatewayFastpathCapability().withProperties(new VirtualNetworkApplianceCapabilityProperties()
                .withIpVersion(VirtualNetworkApplianceCapabilityIpVersion.DUAL_STACK)),
            com.azure.core.util.Context.NONE);
    }
}
