import com.azure.core.util.Context;

/** Samples for AvailabilityGroupListeners Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/sqlvirtualmachine/resource-manager/Microsoft.SqlVirtualMachine/preview/2022-07-01-preview/examples/GetAvailabilityGroupListener.json
     */
    /**
     * Sample code: Gets an availability group listener.
     *
     * @param manager Entry point to SqlVirtualMachineManager.
     */
    public static void getsAnAvailabilityGroupListener(
        com.azure.resourcemanager.sqlvirtualmachine.SqlVirtualMachineManager manager) {
        manager.availabilityGroupListeners().getWithResponse("testrg", "testvmgroup", "agl-test", null, Context.NONE);
    }
}
