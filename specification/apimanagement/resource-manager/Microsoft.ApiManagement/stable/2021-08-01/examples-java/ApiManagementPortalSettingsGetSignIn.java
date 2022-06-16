import com.azure.core.util.Context;

/** Samples for SignInSettings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementPortalSettingsGetSignIn.json
     */
    /**
     * Sample code: ApiManagementPortalSettingsGetSignIn.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementPortalSettingsGetSignIn(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.signInSettings().getWithResponse("rg1", "apimService1", Context.NONE);
    }
}
