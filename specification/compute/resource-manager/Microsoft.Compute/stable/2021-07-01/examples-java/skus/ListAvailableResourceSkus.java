import com.azure.core.util.Context;

/** Samples for ResourceSkus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/skus/ListAvailableResourceSkus.json
     */
    /**
     * Sample code: Lists all available Resource SKUs.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listsAllAvailableResourceSKUs(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getResourceSkus().list(null, null, Context.NONE);
    }
}
