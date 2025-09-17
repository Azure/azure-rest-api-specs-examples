
import com.azure.resourcemanager.appcontainers.models.AppRegistration;
import com.azure.resourcemanager.appcontainers.models.AuthPlatform;
import com.azure.resourcemanager.appcontainers.models.BlobStorageTokenStore;
import com.azure.resourcemanager.appcontainers.models.EncryptionSettings;
import com.azure.resourcemanager.appcontainers.models.Facebook;
import com.azure.resourcemanager.appcontainers.models.GlobalValidation;
import com.azure.resourcemanager.appcontainers.models.IdentityProviders;
import com.azure.resourcemanager.appcontainers.models.Login;
import com.azure.resourcemanager.appcontainers.models.TokenStore;
import com.azure.resourcemanager.appcontainers.models.UnauthenticatedClientActionV2;

/**
 * Samples for ContainerAppsAuthConfigs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * AuthConfigs_BlobStorageTokenStore_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or Update Container App AuthConfig with msi managedIdentityResourceId blob storage token
     * store.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void createOrUpdateContainerAppAuthConfigWithMsiManagedIdentityResourceIdBlobStorageTokenStore(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsAuthConfigs().define("current").withExistingContainerApp("rg1", "myapp")
            .withPlatform(new AuthPlatform().withEnabled(true))
            .withGlobalValidation(
                new GlobalValidation().withUnauthenticatedClientAction(UnauthenticatedClientActionV2.ALLOW_ANONYMOUS))
            .withIdentityProviders(new IdentityProviders().withFacebook(new Facebook().withRegistration(
                new AppRegistration().withAppId("123").withAppSecretSettingName("fakeTokenPlaceholder"))))
            .withLogin(new Login().withTokenStore(new TokenStore().withAzureBlobStorage(new BlobStorageTokenStore()
                .withBlobContainerUri("https://test.blob.core.windows.net/container1").withManagedIdentityResourceId(
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity1"))))
            .withEncryptionSettings(
                new EncryptionSettings().withContainerAppAuthEncryptionSecretName("fakeTokenPlaceholder")
                    .withContainerAppAuthSigningSecretName("fakeTokenPlaceholder"))
            .create();
    }
}
