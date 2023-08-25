import com.azure.resourcemanager.apimanagement.fluent.models.PolicyContractInner;
import com.azure.resourcemanager.apimanagement.models.PolicyContentFormat;
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/** Samples for ApiPolicy CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiPolicy.json
     */
    /**
     * Sample code: ApiManagementCreateApiPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiPolicies()
            .createOrUpdateWithResponse(
                "rg1",
                "apimService1",
                "5600b57e7e8880006a040001",
                PolicyIdName.POLICY,
                new PolicyContractInner()
                    .withValue(
                        "<policies> <inbound /> <backend>    <forward-request />  </backend>  <outbound /></policies>")
                    .withFormat(PolicyContentFormat.XML),
                "*",
                com.azure.core.util.Context.NONE);
    }
}
