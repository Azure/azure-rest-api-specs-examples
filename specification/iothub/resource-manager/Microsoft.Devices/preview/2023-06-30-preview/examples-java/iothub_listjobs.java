/** Samples for IotHubResource ListJobs. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2023-06-30-preview/examples/iothub_listjobs.json
     */
    /**
     * Sample code: IotHubResource_ListJobs.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceListJobs(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().listJobs("myResourceGroup", "testHub", com.azure.core.util.Context.NONE);
    }
}
