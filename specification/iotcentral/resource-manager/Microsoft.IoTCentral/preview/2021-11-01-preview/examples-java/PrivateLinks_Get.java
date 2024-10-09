
/**
 * Samples for PrivateLinks Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/
     * PrivateLinks_Get.json
     */
    /**
     * Sample code: PrivateLinks_Get.
     * 
     * @param manager Entry point to IotCentralManager.
     */
    public static void privateLinksGet(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.privateLinks().getWithResponse("resRg", "myIoTCentralApp", "iotApp", com.azure.core.util.Context.NONE);
    }
}
