
/**
 * Samples for AzureKeyVaultSecretProviderClasses List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-08-21-preview/AzureKeyVaultSecretProviderClasses_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: AzureKeyVaultSecretProviderClasses_ListBySubscription.
     * 
     * @param manager Entry point to SecretsStoreExtensionManager.
     */
    public static void azureKeyVaultSecretProviderClassesListBySubscription(
        com.azure.resourcemanager.secretsstoreextension.SecretsStoreExtensionManager manager) {
        manager.azureKeyVaultSecretProviderClasses().list(com.azure.core.util.Context.NONE);
    }
}
