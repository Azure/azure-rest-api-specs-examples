/** Samples for Products GetProducts. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestack/resource-manager/Microsoft.AzureStack/stable/2022-06-01/examples/Product/ListPost.json
     */
    /**
     * Sample code: Returns a list of products.
     *
     * @param manager Entry point to AzureStackManager.
     */
    public static void returnsAListOfProducts(com.azure.resourcemanager.azurestack.AzureStackManager manager) {
        manager
            .products()
            .getProductsWithResponse("azurestack", "testregistration", "_all", null, com.azure.core.util.Context.NONE);
    }
}
