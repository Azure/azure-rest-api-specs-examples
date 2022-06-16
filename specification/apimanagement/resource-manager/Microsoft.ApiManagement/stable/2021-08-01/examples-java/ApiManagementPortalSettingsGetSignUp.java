import com.azure.core.util.Context;

/** Samples for SignUpSettings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementPortalSettingsGetSignUp.json
     */
    /**
     * Sample code: ApiManagementPortalSettingsGetSignUp.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementPortalSettingsGetSignUp(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.signUpSettings().getWithResponse("rg1", "apimService1", Context.NONE);
    }
}
