/** Samples for Projects List. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/Projects_List.json
     */
    /**
     * Sample code: Projects_List.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void projectsList(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.projects().list("examples-rg", "examples-storageMoverName", com.azure.core.util.Context.NONE);
    }
}
