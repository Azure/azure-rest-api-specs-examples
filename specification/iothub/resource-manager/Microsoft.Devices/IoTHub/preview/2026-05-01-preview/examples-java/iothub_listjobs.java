
/**
 * Samples for IotHubResource ListJobs.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_listjobs.json
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
