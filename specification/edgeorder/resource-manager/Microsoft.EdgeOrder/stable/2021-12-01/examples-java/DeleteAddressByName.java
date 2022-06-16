import com.azure.core.util.Context;

/** Samples for ResourceProvider Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/DeleteAddressByName.json
     */
    /**
     * Sample code: DeleteAddressByName.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void deleteAddressByName(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().delete("TestRG", "TestAddressName1", Context.NONE);
    }
}
