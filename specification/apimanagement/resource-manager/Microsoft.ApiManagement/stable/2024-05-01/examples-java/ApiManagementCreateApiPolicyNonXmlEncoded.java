
import com.azure.resourcemanager.apimanagement.fluent.models.PolicyContractInner;
import com.azure.resourcemanager.apimanagement.models.PolicyContentFormat;
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for ApiPolicy CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateApiPolicyNonXmlEncoded.json
     */
    /**
     * Sample code: ApiManagementCreateApiPolicyNonXmlEncoded.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiPolicyNonXmlEncoded(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiPolicies().createOrUpdateWithResponse("rg1", "apimService1", "5600b57e7e8880006a040001",
            PolicyIdName.POLICY,
            new PolicyContractInner().withValue(
                "<policies>\r\n     <inbound>\r\n     <base />\r\n  <set-header name=\"newvalue\" exists-action=\"override\">\r\n   <value>\"@(context.Request.Headers.FirstOrDefault(h => h.Ke==\"Via\"))\" </value>\r\n    </set-header>\r\n  </inbound>\r\n      </policies>")
                .withFormat(PolicyContentFormat.RAWXML),
            "*", com.azure.core.util.Context.NONE);
    }
}
