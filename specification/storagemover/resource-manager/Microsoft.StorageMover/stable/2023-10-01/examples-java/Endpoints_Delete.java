/** Samples for Endpoints Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-10-01/examples/Endpoints_Delete.json
     */
    /**
     * Sample code: Endpoints_Delete.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void endpointsDelete(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .endpoints()
            .delete(
                "examples-rg", "examples-storageMoverName", "examples-endpointName", com.azure.core.util.Context.NONE);
    }
}
