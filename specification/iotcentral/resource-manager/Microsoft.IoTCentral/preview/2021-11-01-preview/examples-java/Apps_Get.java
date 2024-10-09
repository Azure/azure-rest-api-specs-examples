
/**
 * Samples for Apps GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_Get.json
     */
    /**
     * Sample code: Apps_Get.
     * 
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsGet(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.apps().getByResourceGroupWithResponse("resRg", "myIoTCentralApp", com.azure.core.util.Context.NONE);
    }
}
