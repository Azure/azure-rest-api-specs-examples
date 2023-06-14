/** Samples for AvailabilityGroupListeners Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-08-01-preview/examples/DeleteAvailabilityGroupListener.json
     */
    /**
     * Sample code: Deletes an availability group listener.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void deletesAnAvailabilityGroupListener(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager
            .availabilityGroupListeners()
            .delete("testrg", "testvmgroup", "agl-test", com.azure.core.util.Context.NONE);
    }
}
