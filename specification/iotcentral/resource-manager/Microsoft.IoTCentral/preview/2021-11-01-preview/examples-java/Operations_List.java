import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void operationsList(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.operations().list(Context.NONE);
    }
}
