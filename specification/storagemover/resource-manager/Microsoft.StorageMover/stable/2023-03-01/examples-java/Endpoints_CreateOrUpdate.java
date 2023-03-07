import com.azure.resourcemanager.storagemover.models.EndpointBaseProperties;

/** Samples for Endpoints CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-03-01/examples/Endpoints_CreateOrUpdate.json
     */
    /**
     * Sample code: Endpoints_CreateOrUpdate.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void endpointsCreateOrUpdate(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .endpoints()
            .define("examples-endpointName")
            .withExistingStorageMover("examples-rg", "examples-storageMoverName")
            .withProperties((EndpointBaseProperties) null)
            .create();
    }
}
