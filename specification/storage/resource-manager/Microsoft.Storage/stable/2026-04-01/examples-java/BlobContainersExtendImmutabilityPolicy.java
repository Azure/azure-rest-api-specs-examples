
import com.azure.resourcemanager.storage.fluent.models.ImmutabilityPolicyInner;

/**
 * Samples for BlobContainers ExtendImmutabilityPolicy.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/BlobContainersExtendImmutabilityPolicy.json
     */
    /**
     * Sample code: ExtendImmutabilityPolicy.
     * 
     * @param manager Entry point to StorageManager.
     */
    public static void extendImmutabilityPolicy(com.azure.resourcemanager.storage.StorageManager manager) {
        manager.serviceClient().getBlobContainers().extendImmutabilityPolicyWithResponse("res6238", "sto232",
            "container5023", "8d59f830d0c3bf9",
            new ImmutabilityPolicyInner().withImmutabilityPeriodSinceCreationInDays(100),
            com.azure.core.util.Context.NONE);
    }
}
