/** Samples for IotHubResource GetJob. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2023-06-30-preview/examples/iothub_getjob.json
     */
    /**
     * Sample code: IotHubResource_GetJob.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGetJob(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubResources()
            .getJobWithResponse("myResourceGroup", "testHub", "test", com.azure.core.util.Context.NONE);
    }
}
