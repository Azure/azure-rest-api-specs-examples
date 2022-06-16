import com.azure.core.util.Context;

/** Samples for ResourceProvider ListProductFamiliesMetadata. */
public final class Main {
    /*
     * x-ms-original-file: specification/edgeorder/resource-manager/Microsoft.EdgeOrder/stable/2021-12-01/examples/ListProductFamiliesMetadata.json
     */
    /**
     * Sample code: ListProductFamiliesMetadata.
     *
     * @param manager Entry point to EdgeOrderManager.
     */
    public static void listProductFamiliesMetadata(com.azure.resourcemanager.edgeorder.EdgeOrderManager manager) {
        manager.resourceProviders().listProductFamiliesMetadata(null, Context.NONE);
    }
}
