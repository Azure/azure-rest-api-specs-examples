import com.azure.core.util.Context;

/** Samples for Images Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/Images_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Images_Delete_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void imagesDeleteMaximumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getImages()
            .delete("rgcompute", "aaaaaaaaaaaaaaaaaaaaaaaaaaa", Context.NONE);
    }
}
