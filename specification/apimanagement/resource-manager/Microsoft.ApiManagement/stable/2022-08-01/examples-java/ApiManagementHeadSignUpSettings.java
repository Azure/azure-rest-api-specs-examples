/** Samples for SignUpSettings GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadSignUpSettings.json
     */
    /**
     * Sample code: ApiManagementHeadSignUpSettings.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadSignUpSettings(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.signUpSettings().getEntityTagWithResponse("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
