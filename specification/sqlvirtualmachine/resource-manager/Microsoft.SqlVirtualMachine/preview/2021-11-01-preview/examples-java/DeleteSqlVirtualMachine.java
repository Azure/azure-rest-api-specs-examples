import com.azure.core.util.Context;

/** Samples for SqlVirtualMachines Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/DeleteSqlVirtualMachine.json
     */
    /**
     * Sample code: Deletes a SQL virtual machine.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void deletesASQLVirtualMachine(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().delete("testrg", "testvm1", Context.NONE);
    }
}
