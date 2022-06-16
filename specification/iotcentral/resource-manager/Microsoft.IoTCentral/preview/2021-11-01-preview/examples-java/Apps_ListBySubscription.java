import com.azure.core.util.Context;

/** Samples for Apps List. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_ListBySubscription.json
     */
    /**
     * Sample code: Apps_ListBySubscription.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsListBySubscription(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.apps().list(Context.NONE);
    }
}
