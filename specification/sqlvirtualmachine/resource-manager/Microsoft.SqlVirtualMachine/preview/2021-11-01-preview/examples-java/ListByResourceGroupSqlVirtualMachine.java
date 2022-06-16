import com.azure.core.util.Context;

/** Samples for SqlVirtualMachines ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/ListByResourceGroupSqlVirtualMachine.json
     */
    /**
     * Sample code: Gets all SQL virtual machines in a resource group.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void getsAllSQLVirtualMachinesInAResourceGroup(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().listByResourceGroup("testrg", Context.NONE);
    }
}
