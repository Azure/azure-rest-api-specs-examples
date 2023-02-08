import com.azure.resourcemanager.workloads.models.StopRequest;

/** Samples for SapCentralInstances StopInstance. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPCentralInstances_StopInstance.json
     */
    /**
     * Sample code: Stop the SAP Central Services Instance.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void stopTheSAPCentralServicesInstance(com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapCentralInstances()
            .stopInstance(
                "test-rg",
                "X00",
                "centralServer",
                new StopRequest().withSoftStopTimeoutSeconds(1200L),
                com.azure.core.util.Context.NONE);
    }
}
