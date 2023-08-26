import com.azure.resourcemanager.apimanagement.models.ContentFormat;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiUsingSwaggerImport.json
     */
    /**
     * Sample code: ApiManagementCreateApiUsingSwaggerImport.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiUsingSwaggerImport(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("petstore")
            .withExistingService("rg1", "apimService1")
            .withValue("http://petstore.swagger.io/v2/swagger.json")
            .withFormat(ContentFormat.SWAGGER_LINK_JSON)
            .withPath("petstore")
            .create();
    }
}
