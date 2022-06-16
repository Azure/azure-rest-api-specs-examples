import com.azure.core.util.Context;

/** Samples for SqlVirtualMachineGroups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/ListSubscriptionSqlVirtualMachineGroup.json
     */
    /**
     * Sample code: Gets all SQL virtual machine groups in a subscription.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void getsAllSQLVirtualMachineGroupsInASubscription(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.sqlVirtualMachineGroups().list(Context.NONE);
    }
}
