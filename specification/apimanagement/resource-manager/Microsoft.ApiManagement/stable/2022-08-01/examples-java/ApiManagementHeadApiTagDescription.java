/** Samples for ApiTagDescription GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadApiTagDescription.json
     */
    /**
     * Sample code: ApiManagementHeadApiTagDescription.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApiTagDescription(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiTagDescriptions()
            .getEntityTagWithResponse(
                "rg1",
                "apimService1",
                "59d6bb8f1f7fab13dc67ec9b",
                "59306a29e4bbd510dc24e5f9",
                com.azure.core.util.Context.NONE);
    }
}
