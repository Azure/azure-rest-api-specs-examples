import com.azure.core.util.Context;

/** Samples for GuestAgents ListByVm. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GuestAgent_ListByVm.json
     */
    /**
     * Sample code: GuestAgentListByVm.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void guestAgentListByVm(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.guestAgents().listByVm("testrg", "ContosoVm", Context.NONE);
    }
}
