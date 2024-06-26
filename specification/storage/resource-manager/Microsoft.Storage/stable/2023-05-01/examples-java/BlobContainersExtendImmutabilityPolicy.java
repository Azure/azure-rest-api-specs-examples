
import com.azure.resourcemanager.storage.fluent.models.ImmutabilityPolicyInner;

/**
 * Samples for BlobContainers ExtendImmutabilityPolicy.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/
     * BlobContainersExtendImmutabilityPolicy.json
     */
    /**
     * Sample code: ExtendImmutabilityPolicy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void extendImmutabilityPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.storageAccounts().manager().serviceClient().getBlobContainers().extendImmutabilityPolicyWithResponse(
            "res6238", "sto232", "container5023", "8d59f830d0c3bf9",
            new ImmutabilityPolicyInner().withImmutabilityPeriodSinceCreationInDays(100),
            com.azure.core.util.Context.NONE);
    }
}
