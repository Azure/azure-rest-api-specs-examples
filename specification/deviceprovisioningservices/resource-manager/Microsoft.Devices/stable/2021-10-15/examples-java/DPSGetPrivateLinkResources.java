import com.azure.core.util.Context;

/** Samples for IotDpsResource GetPrivateLinkResources. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceprovisioningservices/resource-manager/Microsoft.Devices/stable/2021-10-15/examples/DPSGetPrivateLinkResources.json
     */
    /**
     * Sample code: PrivateLinkResources_List.
     *
     * @param manager Entry point to IotDpsManager.
     */
    public static void privateLinkResourcesList(
        com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager
            .iotDpsResources()
            .getPrivateLinkResourcesWithResponse(
                "myResourceGroup", "myFirstProvisioningService", "iotDps", Context.NONE);
    }
}
