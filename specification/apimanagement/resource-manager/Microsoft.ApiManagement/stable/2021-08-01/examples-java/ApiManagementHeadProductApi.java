import com.azure.core.util.Context;

/** Samples for ProductApi CheckEntityExists. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadProductApi.json
     */
    /**
     * Sample code: ApiManagementHeadProductApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadProductApi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .productApis()
            .checkEntityExistsWithResponse(
                "rg1", "apimService1", "5931a75ae4bbd512a88c680b", "59306a29e4bbd510dc24e5f9", Context.NONE);
    }
}
