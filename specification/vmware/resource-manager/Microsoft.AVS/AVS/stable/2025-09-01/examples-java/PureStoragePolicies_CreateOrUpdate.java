
import com.azure.resourcemanager.avs.models.PureStoragePolicyProperties;

/**
 * Samples for PureStoragePolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PureStoragePolicies_CreateOrUpdate.json
     */
    /**
     * Sample code: PureStoragePolicies_CreateOrUpdate.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void pureStoragePoliciesCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.pureStoragePolicies().define("storagePolicy1").withExistingPrivateCloud("group1", "cloud1")
            .withProperties(new PureStoragePolicyProperties().withStoragePolicyDefinition("storagePolicyDefinition1")
                .withStoragePoolId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/PureStorage.Block/storagePools/storagePool1"))
            .create();
    }
}
