
/** Samples for DataPolicyManifests GetByPolicyMode. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Authorization/stable/2020-09-01/examples/getDataPolicyManifest
     * .json
     */
    /**
     * Sample code: Retrieve a data policy manifest by policy mode.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void retrieveADataPolicyManifestByPolicyMode(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getDataPolicyManifests()
            .getByPolicyModeWithResponse("Microsoft.KeyVault.Data", com.azure.core.util.Context.NONE);
    }
}
