
import com.azure.resourcemanager.apimanagement.fluent.models.PolicyContractInner;
import com.azure.resourcemanager.apimanagement.models.PolicyContentFormat;
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspaceApiPolicy CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspaceApiPolicy.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspaceApiPolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspaceApiPolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspaceApiPolicies().createOrUpdateWithResponse("rg1", "apimService1", "wks1",
            "5600b57e7e8880006a040001", PolicyIdName.POLICY,
            new PolicyContractInner()
                .withValue(
                    "<policies> <inbound /> <backend>    <forward-request />  </backend>  <outbound /></policies>")
                .withFormat(PolicyContentFormat.XML),
            "*", com.azure.core.util.Context.NONE);
    }
}
