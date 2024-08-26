
/**
 * Samples for Skus ListByOffer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * ListSkusByOffer.json
     */
    /**
     * Sample code: List SKU resources by offer for the HCI Cluster.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listSKUResourcesByOfferForTheHCICluster(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.skus().listByOffer("test-rg", "myCluster", "publisher1", "offer1", null,
            com.azure.core.util.Context.NONE);
    }
}
