import com.azure.resourcemanager.workloads.models.StopRequest;

/** Samples for SapVirtualInstances Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPVirtualInstances_Stop.json
     */
    /**
     * Sample code: SAPVirtualInstances_Stop.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPVirtualInstancesStop(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapVirtualInstances()
            .stop("test-rg", "X00", new StopRequest().withSoftStopTimeoutSeconds(0L), com.azure.core.util.Context.NONE);
    }
}
