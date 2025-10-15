
import com.azure.resourcemanager.storagemover.models.EndpointBaseProperties;

/**
 * Samples for Endpoints CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/Endpoints_CreateOrUpdate_SmbMount.json
     */
    /**
     * Sample code: Endpoints_CreateOrUpdate_SmbMount.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void
        endpointsCreateOrUpdateSmbMount(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.endpoints().define("examples-endpointName")
            .withExistingStorageMover("examples-rg", "examples-storageMoverName")
            .withProperties((EndpointBaseProperties) null).create();
    }
}
