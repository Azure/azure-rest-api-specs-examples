/** Samples for SqlVirtualMachines Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/DeleteSqlVirtualMachine.json
     */
    /**
     * Sample code: Deletes a SQL virtual machine.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void deletesASQLVirtualMachine(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().delete("testrg", "testvm1", com.azure.core.util.Context.NONE);
    }
}
