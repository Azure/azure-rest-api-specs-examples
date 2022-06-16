import com.azure.core.util.Context;

/** Samples for Apps GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/stable/2021-06-01/examples/Apps_Get.json
     */
    /**
     * Sample code: Apps_Get.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsGet(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.apps().getByResourceGroupWithResponse("resRg", "myIoTCentralApp", Context.NONE);
    }
}
