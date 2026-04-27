
/**
 * Samples for Offers ListByPublisher.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/ListOffersByPublisher.json
     */
    /**
     * Sample code: List Offer resources by publisher for the HCI Cluster.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void listOfferResourcesByPublisherForTheHCICluster(
        com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.offers().listByPublisher("test-rg", "myCluster", "publisher1", null, com.azure.core.util.Context.NONE);
    }
}
