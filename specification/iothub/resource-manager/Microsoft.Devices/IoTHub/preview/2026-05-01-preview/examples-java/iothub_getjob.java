
/**
 * Samples for IotHubResource GetJob.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_getjob.json
     */
    /**
     * Sample code: IotHubResource_GetJob.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGetJob(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getJobWithResponse("myResourceGroup", "testHub", "test",
            com.azure.core.util.Context.NONE);
    }
}
