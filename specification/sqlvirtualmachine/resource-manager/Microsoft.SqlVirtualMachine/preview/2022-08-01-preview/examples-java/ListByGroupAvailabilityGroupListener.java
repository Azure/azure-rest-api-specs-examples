/** Samples for AvailabilityGroupListeners ListByGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/ListByGroupAvailabilityGroupListener.json
     */
    /**
     * Sample code: Lists all availability group listeners in a SQL virtual machine group.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void listsAllAvailabilityGroupListenersInASQLVirtualMachineGroup(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.availabilityGroupListeners().listByGroup("testrg", "testvmgroup", com.azure.core.util.Context.NONE);
    }
}
