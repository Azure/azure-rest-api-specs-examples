import com.azure.core.util.Context;

/** Samples for PrivateLinkResources ListByAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/PrivateLinkResources/PrivateLinkResources_ListByAccount.json
     */
    /**
     * Sample code: PrivateLinkResourcesList.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void privateLinkResourcesList(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.privateLinkResources().listByAccount("test-rg", "contoso", Context.NONE);
    }
}
