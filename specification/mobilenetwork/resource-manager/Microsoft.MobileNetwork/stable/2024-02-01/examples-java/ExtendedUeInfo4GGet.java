
/**
 * Samples for ExtendedUeInformation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/
     * ExtendedUeInfo4GGet.json
     */
    /**
     * Sample code: Get UE Information 4G.
     * 
     * @param manager Entry point to MobileNetworkManager.
     */
    public static void getUEInformation4G(com.azure.resourcemanager.mobilenetwork.MobileNetworkManager manager) {
        manager.extendedUeInformations().getWithResponse("rg1", "TestPacketCoreCP", "84449105622",
            com.azure.core.util.Context.NONE);
    }
}
