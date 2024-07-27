
/**
 * Samples for DataPolicyManifests List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2020-09-01/examples/
     * listDataPolicyManifests.json
     */
    /**
     * Sample code: List data policy manifests.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDataPolicyManifests(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().policyClient().getDataPolicyManifests().list(null,
            com.azure.core.util.Context.NONE);
    }
}
