import com.azure.core.util.Context;

/** Samples for GuestAgents Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetGuestAgent.json
     */
    /**
     * Sample code: GetGuestAgent.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getGuestAgent(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.guestAgents().getWithResponse("testrg", "ContosoVm", "default", Context.NONE);
    }
}
