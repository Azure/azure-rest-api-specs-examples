import com.azure.core.util.Context;

/** Samples for SqlVirtualMachines StartAssessment. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/StartAssessmentOnSqlVirtualMachine.json
     */
    /**
     * Sample code: Starts Assessment on SQL virtual machine.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void startsAssessmentOnSQLVirtualMachine(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().startAssessment("testrg", "testvm", Context.NONE);
    }
}
