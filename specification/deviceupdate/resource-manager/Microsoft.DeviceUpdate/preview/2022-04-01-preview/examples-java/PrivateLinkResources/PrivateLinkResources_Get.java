import com.azure.core.util.Context;

/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/PrivateLinkResources/PrivateLinkResources_Get.json
     */
    /**
     * Sample code: PrivateLinkResourcesGet.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void privateLinkResourcesGet(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.privateLinkResources().getWithResponse("test-rg", "contoso", "adu", Context.NONE);
    }
}
