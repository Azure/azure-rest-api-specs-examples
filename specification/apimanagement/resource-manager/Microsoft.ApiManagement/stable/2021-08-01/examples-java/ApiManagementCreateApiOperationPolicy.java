import com.azure.resourcemanager.apimanagement.models.PolicyContentFormat;
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for ApiOperationPolicy CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementCreateApiOperationPolicy.json
     */
    /**
     * Sample code: ApiManagementCreateApiOperationPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiOperationPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiOperationPolicies()
            .define(PolicyIdName.POLICY)
            .withExistingOperation("rg1", "apimService1", "5600b57e7e8880006a040001", "5600b57e7e8880006a080001")
            .withValue("<policies> <inbound /> <backend>    <forward-request />  </backend>  <outbound /></policies>")
            .withFormat(PolicyContentFormat.XML)
            .withIfMatch("*")
            .create();
    }
}
