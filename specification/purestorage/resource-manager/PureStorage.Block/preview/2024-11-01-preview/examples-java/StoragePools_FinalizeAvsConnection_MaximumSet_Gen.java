
import com.azure.resourcemanager.purestorageblock.models.ServiceInitializationInfo;
import com.azure.resourcemanager.purestorageblock.models.StoragePoolFinalizeAvsConnectionPost;

/**
 * Samples for StoragePools FinalizeAvsConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01-preview/StoragePools_FinalizeAvsConnection_MaximumSet_Gen.json
     */
    /**
     * Sample code: StoragePools_FinalizeAvsConnection.
     * 
     * @param manager Entry point to PureStorageBlockManager.
     */
    public static void
        storagePoolsFinalizeAvsConnection(com.azure.resourcemanager.purestorageblock.PureStorageBlockManager manager) {
        manager.storagePools().finalizeAvsConnection("rgpurestorage", "storagePoolname",
            new StoragePoolFinalizeAvsConnectionPost().withServiceInitializationDataEnc("hlgzaxrohv")
                .withServiceInitializationData(new ServiceInitializationInfo().withServiceAccountUsername("axchgm")
                    .withServiceAccountPassword("fakeTokenPlaceholder").withVSphereIp("lhbajnykbznxnxpxozyfdjaciennks")
                    .withVSphereCertificate("s")),
            com.azure.core.util.Context.NONE);
    }
}
