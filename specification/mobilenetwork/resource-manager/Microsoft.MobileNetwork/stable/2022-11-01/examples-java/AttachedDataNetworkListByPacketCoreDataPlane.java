import com.azure.core.util.Context;

/** Samples for AttachedDataNetworks ListByPacketCoreDataPlane. */
public final class Main {
    /*
     * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2022-11-01/examples/AttachedDataNetworkListByPacketCoreDataPlane.json
     */
    /**
     * Sample code: List attached data networks in a data plane.
     *
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void listAttachedDataNetworksInADataPlane(
        com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager
            .attachedDataNetworks()
            .listByPacketCoreDataPlane("rg1", "TestPacketCoreCP", "TestPacketCoreDP", Context.NONE);
    }
}
