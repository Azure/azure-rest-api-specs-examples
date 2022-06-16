import com.azure.core.util.Context;

/** Samples for Apps ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_ListByResourceGroup.json
     */
    /**
     * Sample code: Apps_ListByResourceGroup.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsListByResourceGroup(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.apps().listByResourceGroup("resRg", Context.NONE);
    }
}
