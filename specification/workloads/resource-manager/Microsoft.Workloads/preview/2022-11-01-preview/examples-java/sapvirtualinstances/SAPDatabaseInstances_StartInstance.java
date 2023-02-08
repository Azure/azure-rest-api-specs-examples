/** Samples for SapDatabaseInstances StartInstance. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2022-11-01-preview/examples/sapvirtualinstances/SAPDatabaseInstances_StartInstance.json
     */
    /**
     * Sample code: Start the database instance of the SAP system.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void startTheDatabaseInstanceOfTheSAPSystem(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.sapDatabaseInstances().startInstance("test-rg", "X00", "db0", com.azure.core.util.Context.NONE);
    }
}
