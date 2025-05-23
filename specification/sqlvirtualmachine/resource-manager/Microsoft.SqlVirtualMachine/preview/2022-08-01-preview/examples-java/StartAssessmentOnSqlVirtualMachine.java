
/**
 * Samples for SqlVirtualMachines StartAssessment.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/
     * StartAssessmentOnSqlVirtualMachine.json
     */
    /**
     * Sample code: Starts SQL best practices Assessment on SQL virtual machine.
     * 
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void startsSQLBestPracticesAssessmentOnSQLVirtualMachine(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().startAssessment("testrg", "testvm", com.azure.core.util.Context.NONE);
    }
}
