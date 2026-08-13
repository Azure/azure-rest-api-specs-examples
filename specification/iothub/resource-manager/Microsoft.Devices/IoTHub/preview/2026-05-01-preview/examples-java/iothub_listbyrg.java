
/**
 * Samples for IotHubResource ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_listbyrg.json
     */
    /**
     * Sample code: IotHubResource_ListByResourceGroup.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceListByResourceGroup(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
