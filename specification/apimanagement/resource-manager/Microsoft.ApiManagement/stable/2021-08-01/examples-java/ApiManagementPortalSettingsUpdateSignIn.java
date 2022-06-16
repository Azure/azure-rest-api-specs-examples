import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.fluent.models.PortalSigninSettingsInner;

/** Samples for SignInSettings Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementPortalSettingsUpdateSignIn.json
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
            .updateWithResponse(
                "rg1", "apimService1", "*", new PortalSigninSettingsInner().withEnabled(true), Context.NONE);
    }
}
