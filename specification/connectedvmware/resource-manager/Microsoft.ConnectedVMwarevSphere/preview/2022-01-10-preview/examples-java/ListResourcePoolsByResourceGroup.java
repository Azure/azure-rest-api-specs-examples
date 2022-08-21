import com.azure.core.util.Context;

/** Samples for ResourcePools ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListResourcePoolsByResourceGroup.json
     */
    /**
     * Sample code: ListResourcePoolsByResourceGroup.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listResourcePoolsByResourceGroup(
        com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.resourcePools().listByResourceGroup("testrg", Context.NONE);
    }
}
