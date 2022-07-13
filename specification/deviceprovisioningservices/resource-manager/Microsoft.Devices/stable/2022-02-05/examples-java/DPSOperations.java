import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2022-02-05/examples/DPSOperations.json
     */
    /**
     * Sample code: DPSOperations.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSOperations(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.operations().list(Context.NONE);
    }
}
