
import com.azure.resourcemanager.apimanagement.fluent.models.PolicyFragmentContractInner;
import com.azure.resourcemanager.apimanagement.models.PolicyFragmentContentFormat;

/**
 * Samples for WorkspacePolicyFragment CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateWorkspacePolicyFragment.json
     */
    /**
     * Sample code: ApiManagementCreateWorkspacePolicyFragment.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateWorkspacePolicyFragment(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.workspacePolicyFragments().createOrUpdate("rg1", "apimService1", "wks1", "policyFragment1",
            new PolicyFragmentContractInner()
                .withValue("<fragment><json-to-xml apply=\"always\" consider-accept-header=\"false\" /></fragment>")
                .withDescription("A policy fragment example").withFormat(PolicyFragmentContentFormat.XML),
            null, com.azure.core.util.Context.NONE);
    }
}
