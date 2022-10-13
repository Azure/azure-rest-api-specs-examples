import com.azure.core.util.Context;

/** Samples for StorageSpacesOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/DeleteStorageSpace.json
     */
    /**
     * Sample code: DeleteStorageSpace.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void deleteStorageSpace(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .storageSpacesOperations()
            .deleteByResourceGroupWithResponse("test-arcappliance-resgrp", "test-storage", Context.NONE);
    }
}
