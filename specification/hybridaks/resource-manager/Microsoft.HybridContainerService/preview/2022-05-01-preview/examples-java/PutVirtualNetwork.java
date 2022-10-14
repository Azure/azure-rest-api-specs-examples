import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetworksExtendedLocation;
import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetworksProperties;
import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetworksPropertiesInfraVnetProfile;
import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetworksPropertiesInfraVnetProfileHci;
import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetworksPropertiesVipPoolItem;
import com.azure.resourcemanager.hybridcontainerservice.models.VirtualNetworksPropertiesVmipPoolItem;
import java.util.Arrays;

/** Samples for VirtualNetworksOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/PutVirtualNetwork.json
     */
    /**
     * Sample code: PutVirtualNetwork.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void putVirtualNetwork(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .virtualNetworksOperations()
            .define("test-vnet-static")
            .withRegion("westus")
            .withExistingResourceGroup("test-arcappliance-resgrp")
            .withProperties(
                new VirtualNetworksProperties()
                    .withInfraVnetProfile(
                        new VirtualNetworksPropertiesInfraVnetProfile()
                            .withHci(
                                new VirtualNetworksPropertiesInfraVnetProfileHci()
                                    .withMocGroup("target-group")
                                    .withMocLocation("MocLocation")
                                    .withMocVnetName("test-vnet")))
                    .withVipPool(
                        Arrays
                            .asList(
                                new VirtualNetworksPropertiesVipPoolItem()
                                    .withEndIp("192.168.0.50")
                                    .withStartIp("192.168.0.10")))
                    .withVmipPool(
                        Arrays
                            .asList(
                                new VirtualNetworksPropertiesVmipPoolItem()
                                    .withEndIp("192.168.0.130")
                                    .withStartIp("192.168.0.110"))))
            .withExtendedLocation(
                new VirtualNetworksExtendedLocation()
                    .withType("CustomLocation")
                    .withName(
                        "/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"))
            .create();
    }
}
