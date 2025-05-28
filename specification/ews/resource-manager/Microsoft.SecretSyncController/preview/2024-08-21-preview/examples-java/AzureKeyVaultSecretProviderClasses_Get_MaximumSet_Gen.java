
/**
 * Samples for AzureKeyVaultSecretProviderClasses GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-21-preview/AzureKeyVaultSecretProviderClasses_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: AzureKeyVaultSecretProviderClasses_Get.
     * 
     * @param manager Entry point to SecretsStoreExtensionManager.
     */
    public static void azureKeyVaultSecretProviderClassesGet(
        com.azure.resourcemanager.secretsstoreextension.SecretsStoreExtensionManager manager) {
        manager.azureKeyVaultSecretProviderClasses().getByResourceGroupWithResponse("rg-ssc-example",
            "akvspc-ssc-example", com.azure.core.util.Context.NONE);
    }
}
