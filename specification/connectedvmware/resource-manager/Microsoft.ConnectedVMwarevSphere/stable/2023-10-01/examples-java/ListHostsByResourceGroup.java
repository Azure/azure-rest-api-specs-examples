/** Samples for Hosts ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/ListHostsByResourceGroup.json
     */
    /**
     * Sample code: ListHostsByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listHostsByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.hosts().listByResourceGroup("testrg", com.azure.core.util.Context.NONE);
    }
}
