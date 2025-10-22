
/**
 * Samples for IotDpsResource GetPrivateLinkResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-02-01-preview/DPSGetPrivateLinkResources.json
     */
    /**
     * Sample code: PrivateLinkResources_List.
     * 
     * @param manager Entry point to IotDpsManager.
     */
    public static void
        privateLinkResourcesList(com.azure.resourcemanager.deviceprovisioningservices.IotDpsManager manager) {
        manager.iotDpsResources().getPrivateLinkResourcesWithResponse("myResourceGroup", "myFirstProvisioningService",
            "iotDps", com.azure.core.util.Context.NONE);
    }
}
