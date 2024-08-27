
/**
 * Samples for Offers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/GetOffer
     * .json
     */
    /**
     * Sample code: Get Offer.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getOffer(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.offers().getWithResponse("test-rg", "myCluster", "publisher1", "offer1", null,
            com.azure.core.util.Context.NONE);
    }
}
