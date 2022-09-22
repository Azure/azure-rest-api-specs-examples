import com.azure.core.util.Context;

/** Samples for SqlVirtualMachineGroups GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/GetSqlVirtualMachineGroup.json
     */
    /**
     * Sample code: Gets a SQL virtual machine group.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void getsASQLVirtualMachineGroup(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachineGroups().getByResourceGroupWithResponse("testrg", "testvmgroup", Context.NONE);
    }
}
