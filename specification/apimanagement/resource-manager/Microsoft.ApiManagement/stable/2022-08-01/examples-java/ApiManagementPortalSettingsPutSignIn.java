import com.azure.resourcemanager.apimanagement.fluent.models.PortalSigninSettingsInner;

/** Samples for SignInSettings CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPortalSettingsPutSignIn.json
     */
    /**
     * Sample code: ApiManagementPortalSettingsUpdateSignIn.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementPortalSettingsUpdateSignIn(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .signInSettings()
            .createOrUpdateWithResponse(
                "rg1",
                "apimService1",
                new PortalSigninSettingsInner().withEnabled(true),
                "*",
                com.azure.core.util.Context.NONE);
    }
}
