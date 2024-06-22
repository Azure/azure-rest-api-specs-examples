
/**
 * Samples for UeInformation List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/UeInfoList.json
     */
    /**
     * Sample code: Get UE Information.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getUEInformation(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.ueInformations().list("rg1", "TestPacketCoreCP", com.azure.core.util.Context.NONE);
    }
}
