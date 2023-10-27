/** Samples for Hosts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/GetHost.json
     */
    /**
     * Sample code: GetHost.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getHost(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.hosts().getByResourceGroupWithResponse("testrg", "HRHost", com.azure.core.util.Context.NONE);
    }
}
