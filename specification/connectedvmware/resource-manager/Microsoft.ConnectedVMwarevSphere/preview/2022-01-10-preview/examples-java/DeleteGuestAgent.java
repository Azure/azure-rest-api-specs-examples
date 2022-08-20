import com.azure.core.util.Context;

/** Samples for GuestAgents Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteGuestAgent.json
     */
    /**
     * Sample code: DeleteGuestAgent.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteGuestAgent(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.guestAgents().delete("testrg", "ContosoVm", "default", Context.NONE);
    }
}
