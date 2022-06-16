import com.azure.core.util.Context;

/** Samples for Clusters CreateIdentity. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/stable/2022-05-01/examples/CreateClusterIdentity.json
     */
    /**
     * Sample code: Create cluster Identity.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createClusterIdentity(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.clusters().createIdentity("test-rg", "myCluster", Context.NONE);
    }
}
