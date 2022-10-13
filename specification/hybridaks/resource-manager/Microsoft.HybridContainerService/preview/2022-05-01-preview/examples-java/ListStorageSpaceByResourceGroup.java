import com.azure.core.util.Context;

/** Samples for StorageSpacesOperation ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/ListStorageSpaceByResourceGroup.json
     */
    /**
     * Sample code: ListStorageSpaceByResourceGroup.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void listStorageSpaceByResourceGroup(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager.storageSpacesOperations().listByResourceGroup("test-arcappliance-resgrp", Context.NONE);
    }
}
