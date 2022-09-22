import com.azure.core.util.Context;

/** Samples for SqlVirtualMachines GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/GetSqlVirtualMachine.json
     */
    /**
     * Sample code: Gets a SQL virtual machine.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void getsASQLVirtualMachine(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().getByResourceGroupWithResponse("testrg", "testvm", null, Context.NONE);
    }
}
