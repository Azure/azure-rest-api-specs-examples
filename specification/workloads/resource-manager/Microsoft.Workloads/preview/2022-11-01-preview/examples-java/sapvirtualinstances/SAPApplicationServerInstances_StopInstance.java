import com.azure.resourcemanager.workloads.models.StopRequest;

/** Samples for SapApplicationServerInstances StopInstance. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPApplicationServerInstances_StopInstance.json
     */
    /**
     * Sample code: Stop the SAP Application Server Instance.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void stopTheSAPApplicationServerInstance(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapApplicationServerInstances()
            .stopInstance(
                "test-rg",
                "X00",
                "app01",
                new StopRequest().withSoftStopTimeoutSeconds(0L),
                com.azure.core.util.Context.NONE);
    }
}
