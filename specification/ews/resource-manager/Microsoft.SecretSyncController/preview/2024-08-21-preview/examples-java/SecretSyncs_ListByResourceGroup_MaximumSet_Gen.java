
/**
 * Samples for SecretSyncs ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-21-preview/SecretSyncs_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: SecretSyncs_ListByResourceGroup.
     * 
     * @param manager Entry point to SecretsStoreExtensionManager.
     */
    public static void secretSyncsListByResourceGroup(
        com.azure.resourcemanager.secretsstoreextension.SecretsStoreExtensionManager manager) {
        manager.secretSyncs().listByResourceGroup("rg-ssc-example", com.azure.core.util.Context.NONE);
    }
}
