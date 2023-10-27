/** Samples for Hosts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/ListHosts.json
     */
    /**
     * Sample code: ListHosts.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listHosts(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.hosts().list(com.azure.core.util.Context.NONE);
    }
}
