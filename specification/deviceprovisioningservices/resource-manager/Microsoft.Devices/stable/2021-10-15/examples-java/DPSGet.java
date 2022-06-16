import com.azure.core.util.Context;

/** Samples for IotDpsResource GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSGet.json
     */
    /**
     * Sample code: DPSGet.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void dPSGet(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .iotDpsResources()
            .getByResourceGroupWithResponse("myResourceGroup", "myFirstProvisioningService", Context.NONE);
    }
}
