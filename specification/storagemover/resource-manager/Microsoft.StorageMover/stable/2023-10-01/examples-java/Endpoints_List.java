/** Samples for Endpoints List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/Endpoints_List.json
     */
    /**
     * Sample code: Endpoints_List.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void endpointsList(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.endpoints().list("examples-rg", "examples-storageMoverName", com.azure.core.util.Context.NONE);
    }
}
