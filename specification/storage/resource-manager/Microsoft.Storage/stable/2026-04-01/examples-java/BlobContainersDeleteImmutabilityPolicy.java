
/**
 * Samples for BlobContainers DeleteImmutabilityPolicy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/BlobContainersDeleteImmutabilityPolicy.json
     */
    /**
     * Sample code: DeleteImmutabilityPolicy.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void deleteImmutabilityPolicy(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().deleteImmutabilityPolicyWithResponse("res1581", "sto9621",
            "container4910", "8d59f81a7fa7be0", com.azure.core.util.Context.NONE);
    }
}
