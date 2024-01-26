
/** Samples for DataPolicyManifests List. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-09-01/examples/
     * listDataPolicyManifestsNamespaceFilter.json
     */
    /**
     * Sample code: List data policy manifests with namespace filter.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listDataPolicyManifestsWithNamespaceFilter(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getDataPolicyManifests()
            .list("namespace eq 'Microsoft.KeyVault'", com.azure.core.util.Context.NONE);
    }
}
