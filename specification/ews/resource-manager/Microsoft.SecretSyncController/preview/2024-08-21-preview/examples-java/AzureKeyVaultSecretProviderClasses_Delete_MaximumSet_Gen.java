
/**
 * Samples for AzureKeyVaultSecretProviderClasses Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-21-preview/AzureKeyVaultSecretProviderClasses_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: AzureKeyVaultSecretProviderClasses_Delete.
     * 
     * @param manager Entry point to SecretsStoreExtensionManager.
     */
    public static void azureKeyVaultSecretProviderClassesDelete(
        com.azure.resourcemanager.secretsstoreextension.SecretsStoreExtensionManager manager) {
        manager.azureKeyVaultSecretProviderClasses().delete("rg-ssc-example", "akvspc-ssc-example",
            com.azure.core.util.Context.NONE);
    }
}
