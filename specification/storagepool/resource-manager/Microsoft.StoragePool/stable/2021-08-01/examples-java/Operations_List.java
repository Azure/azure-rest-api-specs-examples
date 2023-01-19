/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagepool/resource-manager/Microsoft.StoragePool/stable/2021-08-01/examples/Operations_List.json
     */
    /**
     * Sample code: List operations.
     *
     * @param manager Entry point to StoragePoolManager.
     */
    public static void listOperations(com.azure.resourcemanager.storagepool.StoragePoolManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
