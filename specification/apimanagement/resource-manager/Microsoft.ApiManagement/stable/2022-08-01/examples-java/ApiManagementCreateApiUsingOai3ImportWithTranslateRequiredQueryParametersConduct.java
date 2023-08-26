import com.azure.resourcemanager.apimanagement.models.ContentFormat;
import com.azure.resourcemanager.apimanagement.models.TranslateRequiredQueryParametersConduct;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiUsingOai3ImportWithTranslateRequiredQueryParametersConduct.json
     */
    /**
     * Sample code: ApiManagementCreateApiUsingOai3ImportWithTranslateRequiredQueryParametersConduct.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiUsingOai3ImportWithTranslateRequiredQueryParametersConduct(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("petstore")
            .withExistingService("rg1", "apimService1")
            .withValue("https://raw.githubusercontent.com/OAI/OpenAPI-Specification/master/examples/v3.0/petstore.yaml")
            .withFormat(ContentFormat.OPENAPI_LINK)
            .withTranslateRequiredQueryParametersConduct(TranslateRequiredQueryParametersConduct.TEMPLATE)
            .withPath("petstore")
            .create();
    }
}
