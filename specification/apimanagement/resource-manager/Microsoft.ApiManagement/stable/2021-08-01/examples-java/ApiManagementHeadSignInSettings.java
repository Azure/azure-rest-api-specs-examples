import com.azure.core.util.Context;

/** Samples for SignInSettings GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadSignInSettings.json
     */
    /**
     * Sample code: ApiManagementHeadSignInSettings.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadSignInSettings(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.signInSettings().getEntityTagWithResponse("rg1", "apimService1", Context.NONE);
    }
}
