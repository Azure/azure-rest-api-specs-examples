import com.azure.core.util.Context;

/** Samples for PrivateLinks List. */
public final class Main {
    /*
     * x-ms-original-file: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateLinks_List.json
     */
    /**
     * Sample code: PrivateLinks_List.
     *
     * @param manager Entry point to IotCentralManager.
     */
    public static void privateLinksList(com.azure.resourcemanager.iotcentral.IotCentralManager manager) {
        manager.privateLinks().list("resRg", "myIoTCentralApp", Context.NONE);
    }
}
