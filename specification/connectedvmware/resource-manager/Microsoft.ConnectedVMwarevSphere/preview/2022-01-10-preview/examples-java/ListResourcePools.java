import com.azure.core.util.Context;

/** Samples for ResourcePools List. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListResourcePools.json
     */
    /**
     * Sample code: ListResourcePools.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listResourcePools(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.resourcePools().list(Context.NONE);
    }
}
