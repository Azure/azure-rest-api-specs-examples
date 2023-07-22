import com.azure.resourcemanager.storagemover.models.EndpointBaseProperties;

/** Samples for Endpoints CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/preview/2023-07-01-preview/examples/Endpoints_CreateOrUpdate_NfsMount.json
     */
    /**
     * Sample code: Endpoints_CreateOrUpdate_NfsMount.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void endpointsCreateOrUpdateNfsMount(
        com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .endpoints()
            .define("examples-endpointName")
            .withExistingStorageMover("examples-rg", "examples-storageMoverName")
            .withProperties((EndpointBaseProperties) null)
            .create();
    }
}
