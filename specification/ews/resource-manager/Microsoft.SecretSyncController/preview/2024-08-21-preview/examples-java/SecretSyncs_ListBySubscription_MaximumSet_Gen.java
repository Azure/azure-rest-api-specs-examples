
/**
 * Samples for SecretSyncs List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-21-preview/SecretSyncs_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: SecretSyncs_ListBySubscription.
     * 
     * @param manager Entry point to SecretsStoreExtensionManager.
     */
    public static void secretSyncsListBySubscription(
        com.azure.resourcemanager.secretsstoreextension.SecretsStoreExtensionManager manager) {
        manager.secretSyncs().list(com.azure.core.util.Context.NONE);
    }
}
