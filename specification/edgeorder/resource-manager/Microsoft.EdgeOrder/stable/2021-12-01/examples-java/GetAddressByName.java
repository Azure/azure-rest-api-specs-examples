
/**
 * Samples for ResourceProvider GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/GetAddressByName.json
     */
    /**
     * Sample code: GetAddressByName.
     * 
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void getAddressByName(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().getByResourceGroupWithResponse("YourResourceGroupName", "TestAddressName1",
            com.azure.core.util.Context.NONE);
    }
}
