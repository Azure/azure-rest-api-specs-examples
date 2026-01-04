
/**
 * Samples for ManagedHsms ListDeleted.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/DeletedManagedHsm_List.json
     */
    /**
     * Sample code: List deleted managed HSMs in the specified subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listDeletedManagedHSMsInTheSpecifiedSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.vaults().manager().serviceClient().getManagedHsms().listDeleted(com.azure.core.util.Context.NONE);
    }
}
