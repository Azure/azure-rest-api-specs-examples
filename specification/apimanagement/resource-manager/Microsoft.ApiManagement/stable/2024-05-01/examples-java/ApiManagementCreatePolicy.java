
import com.azure.resourcemanager.apimanagement.fluent.models.PolicyContractInner;
import com.azure.resourcemanager.apimanagement.models.PolicyContentFormat;
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for Policy CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreatePolicy.json
     */
    /**
     * Sample code: ApiManagementCreatePolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreatePolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.policies().createOrUpdateWithResponse("rg1", "apimService1", PolicyIdName.POLICY,
            new PolicyContractInner().withValue(
                "<policies>\r\n  <inbound />\r\n  <backend>\r\n    <forward-request />\r\n  </backend>\r\n  <outbound />\r\n</policies>")
                .withFormat(PolicyContentFormat.XML),
            null, com.azure.core.util.Context.NONE);
    }
}
