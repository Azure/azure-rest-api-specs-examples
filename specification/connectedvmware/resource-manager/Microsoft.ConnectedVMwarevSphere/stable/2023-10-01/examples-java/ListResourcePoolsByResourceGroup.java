/** Samples for ResourcePools ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/ListResourcePoolsByResourceGroup.json
     */
    /**
     * Sample code: ListResourcePoolsByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listResourcePoolsByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.resourcePools().listByResourceGroup("testrg", com.azure.core.util.Context.NONE);
    }
}
