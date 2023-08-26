import com.azure.resourcemanager.apimanagement.models.ContentFormat;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiUsingImportOverrideServiceUrl.json
     */
    /**
     * Sample code: ApiManagementCreateApiUsingImportOverrideServiceUrl.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiUsingImportOverrideServiceUrl(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("apidocs")
            .withExistingService("rg1", "apimService1")
            .withValue("http://apimpimportviaurl.azurewebsites.net/api/apidocs/")
            .withFormat(ContentFormat.fromString("swagger-link"))
            .withServiceUrl("http://petstore.swagger.wordnik.com/api")
            .withPath("petstoreapi123")
            .create();
    }
}
