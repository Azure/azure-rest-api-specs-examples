import com.azure.core.util.Context;

/** Samples for Hosts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListHosts.json
     */
    /**
     * Sample code: ListHosts.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listHosts(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.hosts().list(Context.NONE);
    }
}
