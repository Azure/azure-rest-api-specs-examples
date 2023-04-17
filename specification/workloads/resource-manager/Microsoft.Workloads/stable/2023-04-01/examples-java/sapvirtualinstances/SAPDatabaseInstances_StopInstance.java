import com.azure.resourcemanager.workloads.models.StopRequest;

/** Samples for SapDatabaseInstances StopInstance. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/stable/2023-04-01/examples/sapvirtualinstances/SAPDatabaseInstances_StopInstance.json
     */
    /**
     * Sample code: Stop the database instance of the SAP system.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void stopTheDatabaseInstanceOfTheSAPSystem(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager
            .sapDatabaseInstances()
            .stopInstance(
                "test-rg",
                "X00",
                "db0",
                new StopRequest().withSoftStopTimeoutSeconds(0L),
                com.azure.core.util.Context.NONE);
    }
}
