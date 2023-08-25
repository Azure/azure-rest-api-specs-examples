/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiRevisionFromExistingApi.json
     */
    /**
     * Sample code: ApiManagementCreateApiRevisionFromExistingApi.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiRevisionFromExistingApi(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("echo-api;rev=3")
            .withExistingService("rg1", "apimService1")
            .withSourceApiId(
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echo-api")
            .withServiceUrl("http://echoapi.cloudapp.net/apiv3")
            .withPath("echo")
            .withApiRevisionDescription("Creating a Revision of an existing API")
            .create();
    }
}
