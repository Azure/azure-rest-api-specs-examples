
/**
 * Samples for Organization ListClusters.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_ClusterList.
     * json
     */
    /**
     * Sample code: Organization_ListClusters.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void organizationListClusters(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().listClusters("myResourceGroup", "myOrganization", "env-12132", 10, null,
            com.azure.core.util.Context.NONE);
    }
}
