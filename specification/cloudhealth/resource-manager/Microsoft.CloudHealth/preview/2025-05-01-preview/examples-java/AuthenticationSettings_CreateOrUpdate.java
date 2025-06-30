
import com.azure.resourcemanager.cloudhealth.models.ManagedIdentityAuthenticationSettingProperties;

/**
 * Samples for AuthenticationSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01-preview/AuthenticationSettings_CreateOrUpdate.json
     */
    /**
     * Sample code: AuthenticationSettings_CreateOrUpdate.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        authenticationSettingsCreateOrUpdate(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.authenticationSettings().define("myAuthSetting")
            .withExistingHealthmodel("myResourceGroup", "myHealthModel")
            .withProperties(new ManagedIdentityAuthenticationSettingProperties().withDisplayName("myDisplayName")
                .withManagedIdentityName("SystemAssigned"))
            .create();
    }
}
