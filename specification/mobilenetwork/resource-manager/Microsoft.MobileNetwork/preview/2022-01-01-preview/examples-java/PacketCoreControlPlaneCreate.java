import com.azure.resourcemanager.mobilenetwork.models.CoreNetworkType;
import com.azure.resourcemanager.mobilenetwork.models.CustomLocationResourceId;
import com.azure.resourcemanager.mobilenetwork.models.InterfaceProperties;
import com.azure.resourcemanager.mobilenetwork.models.MobileNetworkResourceId;

/** Samples for PacketCoreControlPlanes CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-01-01-preview/examples/PacketCoreControlPlaneCreate.json
     */
    /**
     * Sample code: Create packet core control plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createPacketCoreControlPlane(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .packetCoreControlPlanes()
            .define("TestPacketCoreCP")
            .withRegion("eastus")
            .withExistingResourceGroup("rg1")
            .withMobileNetwork(
                new MobileNetworkResourceId()
                    .withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"))
            .withControlPlaneAccessInterface(new InterfaceProperties().withName("N2"))
            .withCustomLocation(
                new CustomLocationResourceId()
                    .withId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ExtendedLocation/customLocations/TestCustomLocation"))
            .withCoreNetworkTechnology(CoreNetworkType.FIVE_GC)
            .withVersion("0.2.0")
            .create();
    }
}
