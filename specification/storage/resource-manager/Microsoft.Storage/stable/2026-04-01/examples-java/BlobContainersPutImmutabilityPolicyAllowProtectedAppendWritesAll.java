
import com.azure.resourcemanager.storage.fluent.models.ImmutabilityPolicyInner;

/**
 * Samples for BlobContainers CreateOrUpdateImmutabilityPolicy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/BlobContainersPutImmutabilityPolicyAllowProtectedAppendWritesAll.json
     */
    /**
     * Sample code: CreateOrUpdateImmutabilityPolicyWithAllowProtectedAppendWritesAll.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void createOrUpdateImmutabilityPolicyWithAllowProtectedAppendWritesAll(
        com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().createOrUpdateImmutabilityPolicyWithResponse(
            "res1782", "sto7069", "container6397", null, new ImmutabilityPolicyInner()
                .withImmutabilityPeriodSinceCreationInDays(3).withAllowProtectedAppendWritesAll(true),
            com.azure.core.util.Context.NONE);
    }
}
