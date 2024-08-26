
/**
 * Samples for Offers ListByCluster.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * ListOffersByCluster.json
     */
    /**
     * Sample code: List Offer resources by HCI Cluster.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listOfferResourcesByHCICluster(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.offers().listByCluster("test-rg", "myCluster", null, com.azure.core.util.Context.NONE);
    }
}
