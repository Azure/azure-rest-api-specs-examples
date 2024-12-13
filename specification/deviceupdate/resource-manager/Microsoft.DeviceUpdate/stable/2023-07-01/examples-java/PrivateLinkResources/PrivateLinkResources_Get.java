
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/
     * PrivateLinkResources/PrivateLinkResources_Get.json
     */
    /**
     * Sample code: PrivateLinkResourcesGet.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void privateLinkResourcesGet(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.privateLinkResources().getWithResponse("test-rg", "contoso", "adu", com.azure.core.util.Context.NONE);
    }
}
