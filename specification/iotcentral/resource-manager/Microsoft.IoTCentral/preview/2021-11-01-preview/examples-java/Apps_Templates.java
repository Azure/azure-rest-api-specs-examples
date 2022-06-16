import com.azure.core.util.Context;

/** Samples for Apps ListTemplates. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/Apps_Templates.json
     */
    /**
     * Sample code: Apps_ListTemplates.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void appsListTemplates(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.apps().listTemplates(Context.NONE);
    }
}
