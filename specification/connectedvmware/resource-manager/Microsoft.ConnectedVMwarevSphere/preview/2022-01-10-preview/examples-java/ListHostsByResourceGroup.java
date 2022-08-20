import com.azure.core.util.Context;

/** Samples for Hosts ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListHostsByResourceGroup.json
     */
    /**
     * Sample code: ListHostsByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listHostsByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.hosts().listByResourceGroup("testrg", Context.NONE);
    }
}
