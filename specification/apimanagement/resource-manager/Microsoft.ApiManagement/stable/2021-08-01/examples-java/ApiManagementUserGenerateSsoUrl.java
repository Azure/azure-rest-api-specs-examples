import com.azure.core.util.Context;

/** Samples for User GenerateSsoUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUserGenerateSsoUrl.json
     */
    /**
     * Sample code: ApiManagementUserGenerateSsoUrl.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUserGenerateSsoUrl(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.users().generateSsoUrlWithResponse("rg1", "apimService1", "57127d485157a511ace86ae7", Context.NONE);
    }
}
