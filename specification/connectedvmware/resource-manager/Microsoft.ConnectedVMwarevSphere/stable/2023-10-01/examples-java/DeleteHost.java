/** Samples for Hosts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/DeleteHost.json
     */
    /**
     * Sample code: DeleteHost.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteHost(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.hosts().delete("testrg", "HRHost", null, com.azure.core.util.Context.NONE);
    }
}
