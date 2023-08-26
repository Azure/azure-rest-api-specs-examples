import com.azure.resourcemanager.apimanagement.models.ApiCreateOrUpdatePropertiesWsdlSelector;
import com.azure.resourcemanager.apimanagement.models.ContentFormat;
import com.azure.resourcemanager.apimanagement.models.SoapApiType;

/** Samples for Api CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateSoapPassThroughApiUsingWsdlImport.json
     */
    /**
     * Sample code: ApiManagementCreateSoapPassThroughApiUsingWsdlImport.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateSoapPassThroughApiUsingWsdlImport(
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
            .withSoapApiType(SoapApiType.SOAP)
            .withPath("currency")
            .create();
    }
}
