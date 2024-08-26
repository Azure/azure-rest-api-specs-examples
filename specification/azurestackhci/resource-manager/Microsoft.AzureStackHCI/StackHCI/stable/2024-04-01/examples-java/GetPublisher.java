
/**
 * Samples for Publishers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/
     * GetPublisher.json
     */
    /**
     * Sample code: Get Publisher.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void getPublisher(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.publishers().getWithResponse("test-rg", "myCluster", "publisher1", com.azure.core.util.Context.NONE);
    }
}
