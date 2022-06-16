import com.azure.core.util.Context;

/** Samples for AvailabilityGroupListeners Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2021-11-01-preview/examples/DeleteAvailabilityGroupListener.json
     */
    /**
     * Sample code: Deletes an availability group listener.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void deletesAnAvailabilityGroupListener(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.availabilityGroupListeners().delete("testrg", "testvmgroup", "agl-test", Context.NONE);
    }
}
