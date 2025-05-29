
/**
 * Samples for SecretSyncs GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-21-preview/SecretSyncs_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: SecretSyncs_Get.
     * 
     * @param manager Entry point to SecretsStoreExtensionManager.
     */
    public static void
        secretSyncsGet(com.azure.resourcemanager.secretsstoreextension.SecretsStoreExtensionManager manager) {
        manager.secretSyncs().getByResourceGroupWithResponse("rg-ssc-example", "secretsync-ssc-example",
            com.azure.core.util.Context.NONE);
    }
}
