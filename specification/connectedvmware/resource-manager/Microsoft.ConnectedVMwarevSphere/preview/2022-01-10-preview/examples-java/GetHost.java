import com.azure.core.util.Context;

/** Samples for Hosts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/GetHost.json
     */
    /**
     * Sample code: GetHost.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void getHost(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.hosts().getByResourceGroupWithResponse("testrg", "HRHost", Context.NONE);
    }
}
