import com.azure.core.util.Context;

/** Samples for Apps Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_Delete.json
     */
    /**
     * Sample code: Apps_Delete.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsDelete(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.apps().delete("resRg", "myIoTCentralApp", Context.NONE);
    }
}
