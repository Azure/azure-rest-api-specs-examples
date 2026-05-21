
import com.azure.resourcemanager.storage.fluent.models.ImmutabilityPolicyInner;

/**
 * Samples for BlobContainers CreateOrUpdateImmutabilityPolicy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01/BlobContainersPutImmutabilityPolicy.json
     */
    /**
     * Sample code: CreateOrUpdateImmutabilityPolicy.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void createOrUpdateImmutabilityPolicy(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().createOrUpdateImmutabilityPolicyWithResponse(
            "res1782", "sto7069", "container6397", null, new ImmutabilityPolicyInner()
                .withImmutabilityPeriodSinceCreationInDays(3).withAllowProtectedAppendWrites(true),
            com.azure.core.util.Context.NONE);
    }
}
