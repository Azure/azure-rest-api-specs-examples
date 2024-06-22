
/**
 * Samples for AttachedDataNetworks Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/
     * AttachedDataNetworkGet.json
     */
    /**
     * Sample code: Get attached data network.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getAttachedDataNetwork(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.attachedDataNetworks().getWithResponse("rg1", "TestPacketCoreCP", "TestPacketCoreDP",
            "TestAttachedDataNetwork", com.azure.core.util.Context.NONE);
    }
}
