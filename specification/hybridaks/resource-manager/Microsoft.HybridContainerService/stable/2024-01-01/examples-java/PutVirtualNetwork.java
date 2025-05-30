
import com.azure.resourcemanager.hybridcontainerservice.models.ExtendedLocationTypes;
import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetworkExtendedLocation;
import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetworkProperties;
import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetworkPropertiesInfraVnetProfile;
import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetworkPropertiesInfraVnetProfileHci;
import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetworkPropertiesVipPoolItem;
import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetworkPropertiesVmipPoolItem;
import java.util.Arrays;

/**
 * Samples for VirtualNetworks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridaks/resource-manager/Microsoft.HybridContainerService/stable/2024-01-01/examples/
     * PutVirtualNetwork.json
     */
    /**
     * Sample code: PutVirtualNetwork.
     * 
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void
        putVirtualNetwork(com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.virtualNetworks().define("test-vnet-static").withRegion("westus")
            .withExistingResourceGroup("test-arcappliance-resgrp")
            .withProperties(new VirtualNetworkProperties()
                .withInfraVnetProfile(new VirtualNetworkPropertiesInfraVnetProfile()
                    .withHci(new VirtualNetworkPropertiesInfraVnetProfileHci().withMocGroup("target-group")
                        .withMocLocation("MocLocation").withMocVnetName("vnet1")))
                .withVipPool(Arrays.asList(
                    new VirtualNetworkPropertiesVipPoolItem().withEndIp("192.168.0.50").withStartIp("192.168.0.10")))
                .withVmipPool(Arrays.asList(
                    new VirtualNetworkPropertiesVmipPoolItem().withEndIp("192.168.0.130").withStartIp("192.168.0.110")))
                .withDnsServers(Arrays.asList("192.168.0.1")).withGateway("192.168.0.1")
                .withIpAddressPrefix("192.168.0.0/16").withVlanId(10))
            .withExtendedLocation(
                new VirtualNetworkExtendedLocation().withType(ExtendedLocationTypes.CUSTOM_LOCATION).withName(
                    "/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"))
            .create();
    }
}
