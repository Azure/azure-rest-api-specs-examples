import com.azure.resourcemanager.apimanagement.models.PolicyFragmentContentFormat;

/** Samples for PolicyFragment CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreatePolicyFragment.json
     */
    /**
     * Sample code: ApiManagementCreatePolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreatePolicy(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .policyFragments()
            .define("policyFragment1")
            .withExistingService("rg1", "apimService1")
            .withValue("<fragment><json-to-xml apply=\"always\" consider-accept-header=\"false\" /></fragment>")
            .withDescription("A policy fragment example")
            .withFormat(PolicyFragmentContentFormat.XML)
            .create();
    }
}
