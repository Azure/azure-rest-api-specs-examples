import com.azure.resourcemanager.hybridcontainerservice.models.StorageSpacesExtendedLocation;
import com.azure.resourcemanager.hybridcontainerservice.models.StorageSpacesProperties;
import com.azure.resourcemanager.hybridcontainerservice.models.StorageSpacesPropertiesHciStorageProfile;

/** Samples for StorageSpacesOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-09-01-preview/examples/PutStorageSpace.json
     */
    /**
     * Sample code: PutStorageSpace.
     *
     * @param manager Entry point to HybridContainerServiceManager.
     */
    public static void putStorageSpace(
        com.azure.resourcemanager.hybridcontainerservice.HybridContainerServiceManager manager) {
        manager
            .storageSpacesOperations()
            .define("test-storage")
            .withRegion("westus")
            .withExistingResourceGroup("test-arcappliance-resgrp")
            .withProperties(
                new StorageSpacesProperties()
                    .withHciStorageProfile(
                        new StorageSpacesPropertiesHciStorageProfile()
                            .withMocGroup("target-group")
                            .withMocLocation("MocLocation")
                            .withMocStorageContainer("WssdStorageContainer")))
            .withExtendedLocation(
                new StorageSpacesExtendedLocation()
                    .withType("CustomLocation")
                    .withName(
                        "/subscriptions/a3e42606-29b1-4d7d-b1d9-9ff6b9d3c71b/resourcegroups/test-arcappliance-resgrp/providers/microsoft.extendedlocation/customlocations/testcustomlocation"))
            .create();
    }
}
