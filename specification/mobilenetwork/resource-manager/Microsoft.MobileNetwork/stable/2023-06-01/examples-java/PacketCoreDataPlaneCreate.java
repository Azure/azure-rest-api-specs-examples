import com.azure.resourcemanager.mobilenetwork.models.InterfaceProperties;

/** Samples for PacketCoreDataPlanes CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-06-01/examples/PacketCoreDataPlaneCreate.json
     */
    /**
     * Sample code: Create packet core data plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void createPacketCoreDataPlane(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .packetCoreDataPlanes()
            .define("testPacketCoreDP")
            .withRegion("eastus")
            .withExistingPacketCoreControlPlane("rg1", "testPacketCoreCP")
            .withUserPlaneAccessInterface(new InterfaceProperties().withName("N3"))
            .create();
    }
}
