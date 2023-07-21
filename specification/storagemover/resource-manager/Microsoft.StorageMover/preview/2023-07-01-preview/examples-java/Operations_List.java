/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void operationsList(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
