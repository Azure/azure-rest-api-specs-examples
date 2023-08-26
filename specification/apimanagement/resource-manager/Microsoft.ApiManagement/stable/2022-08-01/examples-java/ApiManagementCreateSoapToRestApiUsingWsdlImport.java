import com.azure.resourcemanager.apimanagement.models.ApiCreateOrUpdatePropertiesWsdlSelector;
import com.azure.resourcemanager.apimanagement.models.ContentFormat;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateSoapToRestApiUsingWsdlImport.json
     */
    /**
     * Sample code: ApiManagementCreateSoapToRestApiUsingWsdlImport.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateSoapToRestApiUsingWsdlImport(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apis()
            .define("soapApi")
            .withExistingService("rg1", "apimService1")
            .withValue("http://www.webservicex.net/CurrencyConvertor.asmx?WSDL")
            .withFormat(ContentFormat.WSDL_LINK)
            .withWsdlSelector(
                new ApiCreateOrUpdatePropertiesWsdlSelector()
                    .withWsdlServiceName("CurrencyConvertor")
                    .withWsdlEndpointName("CurrencyConvertorSoap"))
            .withPath("currency")
            .create();
    }
}
