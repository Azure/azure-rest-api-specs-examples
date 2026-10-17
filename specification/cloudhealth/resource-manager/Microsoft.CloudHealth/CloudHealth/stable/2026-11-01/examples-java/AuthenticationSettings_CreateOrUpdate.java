
import com.azure.resourcemanager.cloudhealth.models.ManagedIdentityAuthenticationSettingProperties;

/**
 * Samples for AuthenticationSettings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/AuthenticationSettings_CreateOrUpdate.json
     */
    /**
     * Sample code: AuthenticationSettings_CreateOrUpdate.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void
        authenticationSettingsCreateOrUpdate(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.authenticationSettings().define("default-auth")
            .withExistingHealthmodel("online-store-rg", "online-store")
            .withProperties(new ManagedIdentityAuthenticationSettingProperties()
                .withDisplayName("Default managed identity").withManagedIdentityName("SystemAssigned"))
            .create();
    }
}
