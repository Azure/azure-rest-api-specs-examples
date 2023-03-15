/** Samples for StorageSpacesOperation GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/GetStorageSpace.json
     */
    /**
     * Sample code: GetStorageSpace.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void getStorageSpace(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .storageSpacesOperations()
            .getByResourceGroupWithResponse(
                "test-arcappliance-resgrp", "test-storage", com.azure.core.util.Context.NONE);
    }
}
