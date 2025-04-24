
import com.azure.resourcemanager.apimanagement.fluent.models.PolicyContractInner;
import com.azure.resourcemanager.apimanagement.models.PolicyContentFormat;
import com.azure.resourcemanager.apimanagement.models.PolicyIdName;

/**
 * Samples for WorkspacePolicy CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspacePolicy.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspacePolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateWorkspacePolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspacePolicies().createOrUpdateWithResponse("rg1", "apimService1", "wks1", PolicyIdName.POLICY,
            new PolicyContractInner()
                .withValue(
                    "<policies> <inbound /> <backend>    <forward-request />  </backend>  <outbound /></policies>")
                .withFormat(PolicyContentFormat.XML),
            "*", com.azure.core.util.Context.NONE);
    }
}
