import com.azure.resourcemanager.apimanagement.models.ContentFormat;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiUsingOai3Import.json
     */
    /**
     * Sample code: ApiManagementCreateApiUsingOai3Import.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiUsingOai3Import(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("petstore")
            .withExistingService("rg1", "apimService1")
            .withValue("https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/examples/v3.0/petstore.yaml")
            .withFormat(ContentFormat.OPENAPI_LINK)
            .withPath("petstore")
            .create();
    }
}
