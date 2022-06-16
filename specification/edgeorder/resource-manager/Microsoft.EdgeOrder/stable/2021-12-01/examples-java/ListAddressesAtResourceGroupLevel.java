import com.azure.core.util.Context;

/** Samples for ResourceProvider ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListAddressesAtResourceGroupLevel.json
     */
    /**
     * Sample code: ListAddressesAtResourceGroupLevel.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void listAddressesAtResourceGroupLevel(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().listByResourceGroup("TestRG", null, null, Context.NONE);
    }
}
