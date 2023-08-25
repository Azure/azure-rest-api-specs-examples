import com.azure.resourcemanager.apimanagement.models.ContentFormat;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiUsingWadlImport.json
     */
    /**
     * Sample code: ApiManagementCreateApiUsingWadlImport.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiUsingWadlImport(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("petstore")
            .withExistingService("rg1", "apimService1")
            .withValue(
                "https://developer.cisco.com/media/wae-release-6-2-api-reference/wae-collector-rest-api/application.wadl")
            .withFormat(ContentFormat.WADL_LINK_JSON)
            .withPath("collector")
            .create();
    }
}
