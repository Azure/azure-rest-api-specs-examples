import com.azure.core.util.Context;

/** Samples for Hosts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/DeleteHost.json
     */
    /**
     * Sample code: DeleteHost.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void deleteHost(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.hosts().delete("testrg", "HRHost", null, Context.NONE);
    }
}
