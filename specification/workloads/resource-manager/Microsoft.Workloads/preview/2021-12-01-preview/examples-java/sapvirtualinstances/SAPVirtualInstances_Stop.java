import com.azure.core.util.Context;
import com.azure.resourcemanager.workloads.models.StopRequest;

/** Samples for SapVirtualInstances Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Stop.json
     */
    /**
     * Sample code: SAPVirtualInstances_Stop.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void sAPVirtualInstancesStop(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapVirtualInstances().stop("test-rg", "X00", new StopRequest().withHardStop(false), Context.NONE);
    }
}
