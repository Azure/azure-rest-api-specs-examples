
/**
 * Samples for BlobContainers LockImmutabilityPolicy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/BlobContainersLockImmutabilityPolicy.json
     */
    /**
     * Sample code: LockImmutabilityPolicy.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void lockImmutabilityPolicy(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().lockImmutabilityPolicyWithResponse("res2702", "sto5009",
            "container1631", "8d59f825b721dd3", com.azure.core.util.Context.NONE);
    }
}
