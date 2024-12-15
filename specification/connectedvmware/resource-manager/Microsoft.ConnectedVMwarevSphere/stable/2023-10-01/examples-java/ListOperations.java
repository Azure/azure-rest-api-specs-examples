
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/stable/2023-10-01/examples/
     * ListOperations.json
     */
    /**
     * Sample code: ListOperations.
     * 
     * @param manager Entry point to ConnectedVMwareManager.
     */
    public static void listOperations(com.azure.resourcemanager.connectedvmware.ConnectedVMwareManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
