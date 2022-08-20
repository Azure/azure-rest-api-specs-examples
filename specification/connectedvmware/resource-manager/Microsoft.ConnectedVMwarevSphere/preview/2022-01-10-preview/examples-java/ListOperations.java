import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/ListOperations.json
     */
    /**
     * Sample code: ListOperations.
     *
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listOperations(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.operations().list(Context.NONE);
    }
}
