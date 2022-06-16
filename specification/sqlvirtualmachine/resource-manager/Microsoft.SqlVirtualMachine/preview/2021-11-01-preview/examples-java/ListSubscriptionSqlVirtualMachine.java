import com.azure.core.util.Context;

/** Samples for SqlVirtualMachines List. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/ListSubscriptionSqlVirtualMachine.json
     */
    /**
     * Sample code: Gets all SQL virtual machines in a subscription.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void getsAllSQLVirtualMachinesInASubscription(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachines().list(Context.NONE);
    }
}
