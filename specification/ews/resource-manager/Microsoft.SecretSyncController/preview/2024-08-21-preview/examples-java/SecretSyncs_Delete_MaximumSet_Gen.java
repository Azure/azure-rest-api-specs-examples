
/**
 * Samples for SecretSyncs Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-21-preview/SecretSyncs_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: SecretSyncs_Delete.
     * 
     * @param manager Entry point to SecretsStoreExtensionManager.
     */
    public static void
        secretSyncsDelete(com.azure.resourcemanager.secretsstoreextension.SecretsStoreExtensionManager manager) {
        manager.secretSyncs().delete("rg-ssc-example", "secretsync-ssc-example", com.azure.core.util.Context.NONE);
    }
}
