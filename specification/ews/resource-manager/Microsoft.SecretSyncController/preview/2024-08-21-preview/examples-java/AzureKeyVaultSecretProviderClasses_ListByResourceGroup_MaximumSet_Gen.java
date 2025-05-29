
/**
 * Samples for AzureKeyVaultSecretProviderClasses ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-21-preview/AzureKeyVaultSecretProviderClasses_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: AzureKeyVaultSecretProviderClasses_ListByResourceGroup.
     * 
     * @param manager Entry point to SecretsStoreExtensionManager.
     */
    public static void azureKeyVaultSecretProviderClassesListByResourceGroup(
        com.azure.resourcemanager.secretsstoreextension.SecretsStoreExtensionManager manager) {
        manager.azureKeyVaultSecretProviderClasses().listByResourceGroup("rg-ssc-example",
            com.azure.core.util.Context.NONE);
    }
}
