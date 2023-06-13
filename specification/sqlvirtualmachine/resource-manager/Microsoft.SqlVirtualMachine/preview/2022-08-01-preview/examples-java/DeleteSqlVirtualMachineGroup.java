/** Samples for SqlVirtualMachineGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/DeleteSqlVirtualMachineGroup.json
     */
    /**
     * Sample code: Deletes a SQL virtual machine group.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void deletesASQLVirtualMachineGroup(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachineGroups().delete("testrg", "testvmgroup", com.azure.core.util.Context.NONE);
    }
}
