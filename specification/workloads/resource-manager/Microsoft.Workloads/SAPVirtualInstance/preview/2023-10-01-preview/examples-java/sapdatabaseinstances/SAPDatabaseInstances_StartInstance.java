
/**
 * Samples for SapDatabaseInstances StartInstance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/
     * examples/sapdatabaseinstances/SAPDatabaseInstances_StartInstance.json
     */
    /**
     * Sample code: Start the database instance of the SAP system.
     * 
     * @param manager Entry point to WorkloadsSapVirtualInstanceManager.
     */
    public static void startTheDatabaseInstanceOfTheSAPSystem(
        com.azure.resourcemanager.workloadssapvirtualinstance.WorkloadsSapVirtualInstanceManager manager) {
        manager.sapDatabaseInstances().startInstance("test-rg", "X00", "db0", null, com.azure.core.util.Context.NONE);
    }
}
